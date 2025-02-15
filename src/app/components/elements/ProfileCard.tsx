import "@/scss/blog-card.scss";
import React from "react";
import "@/scss/profile.scss";

type ProfileCardProps = {
  name: string;
  email: string;
};

const ProfileCard: React.FC<ProfileCardProps> = ({ name, email }) => {
  return (
    <div className="proflie mt-4">
      <div className="max-w-lg mx-auto bg-beige p-6 rounded-2xl shadow-lg flex flex-col items-center">
        <div className="w-32 h-32 rounded-full overflow-hidden shadow-md">
          <img src="" alt="" className="w-full h-full object-cover" />
        </div>
        <h2 className="text-xl font-bold mt-4 text-gray-800">PROFILE</h2>
        <div className="bg-white rounded-lg p-4 mt-4 w-full shadow-md">
          <p className="text-gray-600 text-sm">Name</p>
          <p className="text-gray-900 font-semibold text-lg">{name}</p>
          <p className="text-gray-600 text-sm mt-2">Email</p>
          <p className="text-gray-900 font-semibold text-lg">{email}</p>
        </div>
      </div>
    </div>
  );
};

export default ProfileCard;
