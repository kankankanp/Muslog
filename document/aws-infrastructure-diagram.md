# AWS Infrastructure Diagram

## Simple Blog Production Environment

```mermaid
graph TB
    Users[Users]
    DNS[Route53 DNS]
    CF[CloudFront Distribution]
    ACM[ACM Certificate]
    
    subgraph VPC["VPC (10.0.0.0/16)"]
        IGW[Internet Gateway]
        ALB[Application Load Balancer]
        
        subgraph PublicSubnets["Public Subnets"]
            PubSub1[Public Subnet AZ-a<br/>10.0.1.0/24]
            PubSub2[Public Subnet AZ-c<br/>10.0.2.0/24]
            NAT1[NAT Gateway AZ-a]
            NAT2[NAT Gateway AZ-c]
        end
        
        subgraph PrivateSubnets["Private Subnets"]
            PrivSub1[Private Subnet AZ-a<br/>10.0.11.0/24]
            PrivSub2[Private Subnet AZ-c<br/>10.0.12.0/24]
            ECS1[ECS Fargate Backend<br/>Port: 8080]
        end
        
        subgraph Database["Aurora PostgreSQL Cluster"]
            RDS1[Aurora Instance AZ-a]
            RDS2[Aurora Instance AZ-c]
        end
    end
    
    subgraph Storage["S3 Storage"]
        S3Frontend[S3 Frontend Bucket]
        S3Media[S3 Media Bucket]
        OAI[Origin Access Identity]
    end
    
    subgraph Container["Container Services"]
        ECR[ECR Backend Repository]
        ECSCluster[ECS Cluster]
        IAMExecRole[IAM Task Execution Role]
    end
    
    %% User Flow
    Users -->|HTTPS Request| DNS
    DNS -->|Route to CloudFront| CF
    
    %% CloudFront Routing
    CF -->|Static Files| S3Frontend
    CF -->|/api/* requests| ALB
    CF -->|Secure Access| OAI
    OAI --> S3Frontend
    OAI --> S3Media
    
    %% Load Balancer Flow
    IGW --> ALB
    ALB -->|Health Check /health| ECS1
    
    %% ECS Configuration
    ECR -->|Container Image| ECS1
    ECSCluster --> ECS1
    IAMExecRole --> ECS1
    
    %% Database Connection
    ECS1 -->|Port 5432| RDS1
    ECS1 -->|Port 5432| RDS2
    
    %% Media Upload
    ECS1 -->|s3:PutObject| S3Media
    
    %% NAT Gateway
    ECS1 --> NAT1
    ECS1 --> NAT2
    NAT1 --> IGW
    NAT2 --> IGW
    
    %% SSL Certificate
    ACM --> CF
    ACM --> ALB
    DNS --> ACM
    
    %% Placement
    PubSub1 --> NAT1
    PubSub2 --> NAT2
    PrivSub1 --> ECS1
    PrivSub2 --> RDS1
    PrivSub2 --> RDS2
    
    %% Styling
    classDef userClass fill:#e1f5fe,stroke:#01579b,stroke-width:2px
    classDef awsGlobal fill:#fff3e0,stroke:#e65100,stroke-width:2px
    classDef networking fill:#f3e5f5,stroke:#4a148c,stroke-width:2px
    classDef compute fill:#e8f5e8,stroke:#1b5e20,stroke-width:2px
    classDef storage fill:#fff8e1,stroke:#e65100,stroke-width:2px
    classDef database fill:#fce4ec,stroke:#880e4f,stroke-width:2px
    
    class Users,DNS userClass
    class CF,ACM awsGlobal
    class IGW,ALB,NAT1,NAT2,PubSub1,PubSub2,PrivSub1,PrivSub2 networking
    class ECS1,ECSCluster,ECR,IAMExecRole compute
    class S3Frontend,S3Media,OAI storage
    class RDS1,RDS2 database
```

## アーキテクチャ概要

### フロントエンド配信
- **S3 + CloudFront**: 静的ファイル（HTML/CSS/JS）の高速配信
- **Route53**: ドメイン名解決とDNS管理
- **ACM**: SSL/TLS証明書による暗号化通信

### バックエンドAPI
- **ECS Fargate**: サーバーレスコンテナでバックエンドAPI実行
- **ALB**: ロードバランシングとHTTPS終端
- **ECR**: Dockerイメージの管理

### データベース
- **Aurora PostgreSQL**: 高可用性クラスター構成
- **マルチAZ**: 冗長性とフェイルオーバー対応

### セキュリティ
- **VPC**: プライベートネットワーク分離
- **セキュリティグループ**: ファイアウォール設定
- **NAT Gateway**: プライベートサブネットからの安全な外部通信

### 主要な特徴
- **完全にサーバーレス**: ECS Fargateによる管理不要なコンテナ実行
- **高可用性**: マルチAZ構成による冗長性
- **セキュア**: プライベートサブネットでの機密データ処理
- **スケーラブル**: CloudFrontとECS Fargateによる自動スケーリング
- **コスト効率**: 使用量ベースの課金によるコスト最適化

### トラフィックフロー
1. ユーザーがドメインにアクセス
2. Route53がCloudFrontにルーティング
3. CloudFrontが静的ファイルはS3から、APIリクエスト（/api/*）はALBから配信
4. ALBがECS Fargateのバックエンドコンテナにルーティング
5. バックエンドがAurora PostgreSQLクラスターにアクセス