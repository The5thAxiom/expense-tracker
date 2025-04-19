import {Outlet, Link} from "react-router-dom"

const Layout = () => {
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
                </ul>
            </nav>

            <Outlet></Outlet>
        </>
    )
    
}

export default Layout;