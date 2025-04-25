import {Outlet, Link} from "react-router-dom"
import LoginButton from "../components/Login"
import LogoutButton from "../components/Logout"
import Profile from "../components/Profile"
import { useAuth0 } from "@auth0/auth0-react"

const Layout = () => {
    const { isAuthenticated } = useAuth0();
    
    return (
        <>
            <nav>
                <ul>
                    <li>
                        <Link to="/">Home</Link>
                    </li>
                    <li>
                        <Link to="/about">About</Link>
                    </li>
                    <li>
                        <Link to="/help">Help</Link>
                    </li>
                    <li>
                        <Link to="/payments">Payment</Link>
                    </li>
                    {!isAuthenticated && 
                    <li><LoginButton></LoginButton></li>}
                    {isAuthenticated && <li><LogoutButton></LogoutButton></li>}
                </ul>
            </nav>
            <Profile></Profile>

            <Outlet></Outlet>
        </>
    )
    
}

export default Layout;