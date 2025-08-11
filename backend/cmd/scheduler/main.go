package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func main() {
	action := os.Getenv("ACTION")
	if action != "start" && action != "stop" {
		log.Fatal("Error: ACTION environment variable must be set to 'start' or 'stop'")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Error loading AWS config: %v", err)
	}

	ecsClient := ecs.NewFromConfig(cfg)
	rdsClient := rds.NewFromConfig(cfg)

	clusterName := os.Getenv("ECS_CLUSTER_NAME")
	serviceName := os.Getenv("ECS_SERVICE_NAME")
	dbClusterIdentifier := os.Getenv("DB_CLUSTER_IDENTIFIER")
	desiredCountStr := os.Getenv("ECS_DESIRED_COUNT")

	if action == "start" {
		log.Println("Starting resources...")
		if err := startRdsCluster(rdsClient, dbClusterIdentifier); err != nil {
			log.Fatalf("Error starting RDS cluster: %v", err)
		}
		// In a real-world scenario, you should poll for the RDS cluster to be available
		// before starting the ECS service. For simplicity, we'll proceed without a wait.
		if err := updateEcsService(ecsClient, clusterName, serviceName, desiredCountStr); err != nil {
			log.Fatalf("Error starting ECS service: %v", err)
		}
		log.Println("Resources started successfully.")
	} else {
		log.Println("Stopping resources...")
		if err := updateEcsService(ecsClient, clusterName, serviceName, "0"); err != nil {
			log.Fatalf("Error stopping ECS service: %v", err)
		}
		if err := stopRdsCluster(rdsClient, dbClusterIdentifier); err != nil {
			log.Fatalf("Error stopping RDS cluster: %v", err)
		}
		log.Println("Resources stopped successfully.")
	}
}

func updateEcsService(client *ecs.Client, clusterName, serviceName, desiredCountStr string) error {
	desiredCount, err := strconv.ParseInt(desiredCountStr, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid desired count: %w", err)
	}
	log.Printf("Updating ECS service %s in cluster %s to desired count %d\n", serviceName, clusterName, desiredCount)
	desiredCount32 := int32(desiredCount)
	_, err = client.UpdateService(context.TODO(), &ecs.UpdateServiceInput{
		Cluster:      &clusterName,
		Service:      &serviceName,
		DesiredCount: &desiredCount32,
	})
	return err
}

func stopRdsCluster(client *rds.Client, dbClusterIdentifier string) error {
	fmt.Printf("Stopping RDS cluster %s\n", dbClusterIdentifier)
	_, err := client.StopDBCluster(context.TODO(), &rds.StopDBClusterInput{
		DBClusterIdentifier: &dbClusterIdentifier,
	})
	return err
}

func startRdsCluster(client *rds.Client, dbClusterIdentifier string) error {
	fmt.Printf("Starting RDS cluster %s\n", dbClusterIdentifier)
	_, err := client.StartDBCluster(context.TODO(), &rds.StartDBClusterInput{
		DBClusterIdentifier: &dbClusterIdentifier,
	})
	return err
}
