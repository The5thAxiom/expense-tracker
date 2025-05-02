import noUserProfileIcon from "../../assets/noUserProfileIcon.png";
import { useAuth0 } from "@auth0/auth0-react";
import "./Profile.css"

const Profile = () => {
  const { user, isAuthenticated, isLoading } = useAuth0();

  if (!isAuthenticated || isLoading) {
    return <img className="profile-icon" src={noUserProfileIcon}></img>;
  } else {
    return <img className="profile-icon" src={user?.picture}></img>;
  }
};

export default Profile;
