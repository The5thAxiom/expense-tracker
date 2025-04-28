import { Outlet } from "react-router-dom";
import Header from "../../components/Header/Header";
import Footer from "../../components/Footer/Footer";
import "./Layout.css";

const Layout = () => {
  return (
    <>
      <div className="layout-div">
        <Header></Header>
        <main>
          <h2> Main content</h2>
          <Outlet></Outlet>
        </main>
        <Footer></Footer>
      </div>
    </>
  );
};

export default Layout;
