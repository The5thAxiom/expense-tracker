import {Link} from "react-router-dom"
import { useAuth0 } from "@auth0/auth0-react";
import LoginButton from "../../components/Login";
import LogoutButton from "../../components/Logout";
import "./Navbar.css"

const Navbar = () => {
  const { isAuthenticated } = useAuth0();

    return (
        <nav className="nav">
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
              <Link to="/payments">Payment </Link>
            </div>
            {!isAuthenticated && (
              <div className="navbar-list-item">
                <LoginButton></LoginButton> 
              </div>
            )}
            {isAuthenticated && (
              <div className="navbar-list-item">
                <LogoutButton></LogoutButton> 
              </div>
            )}
        </nav>
    )
}

export default Navbar;