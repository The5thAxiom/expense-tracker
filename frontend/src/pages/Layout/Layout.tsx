import { Outlet } from "react-router-dom";
import Header from "../../components/Header/Header";

const Layout = () => {

  return (
    <>
      <Header></Header>
      <main>
        <h2> Main content</h2>
        <Outlet></Outlet>
      </main>
      <footer>
        <ul>
          <li>link 1</li>
          <li>link 2</li>
          <li>link 3</li>
        </ul>
      </footer>
    </>
  );
};

export default Layout;
