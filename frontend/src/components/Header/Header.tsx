import { Link } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import LoginButton from "../../components/Login";
import LogoutButton from "../../components/Logout";
import Profile from "../Profile/Profile";
import "./Header.css"
import { useEffect, useState } from "react";

const Header = () => {
  const { isAuthenticated, isLoading, getAccessTokenSilently} = useAuth0();
  const [ token, setToken ] = useState<string>("token");
  
  useEffect(() => {
    async function getToken() {
      if (isAuthenticated && !isLoading) {
        const token = await getAccessTokenSilently();
        setToken(token);
        localStorage.setItem('token', token);
      }
    }
    console.log({isAuthenticated, isLoading})
    getToken();
  }, [isAuthenticated, isLoading]);

  return (
    <header className="header">
      <div className="header-title" onClick={async () => {
        const res = await fetch("/api/test", {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        });
        const data = await res.json();
        console.log(data);
      }}>Expense Tracker</div>
      <nav className="nav-left">
        <div className="navbar-list-item">
          <Link to="/">Home </Link>
        </div>
        <div className="navbar-list-item">
          <Link to="/about">About </Link>
        </div>
        <div className="navbar-list-item">
          <Link to="/help">Help </Link>
        </div>
        <div className="navbar-list-item">
          <Link to="/payments">Expense </Link>
        </div>
        
      </nav>
      <nav className="nav-right">
        <div>
            Settings
        </div>
        {!isAuthenticated && !isLoading &&(
          <div className="navbar-list-item">
            <LoginButton></LoginButton>
          </div>
        )}
        {isAuthenticated && !isLoading && (
          <div className="navbar-list-item">
            <LogoutButton></LogoutButton>
          </div>
        )}
      </nav>
        <Profile></Profile>
    </header>
  );
};

export default Header;
