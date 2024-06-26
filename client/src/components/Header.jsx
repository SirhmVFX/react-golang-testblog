import { Link } from "react-router-dom";

function Header() {
  return (
    <>
      <header className="shadow-gray-100 shadow-md p-4 ">
        <nav className="flex justify-between">
          <h1 className="">A blog</h1>

          <div className="flex gap-4">
            <Link to={"/login"}> Login </Link>
            <Link to={"/signup"}> Signup </Link>
          </div>
        </nav>
      </header>
    </>
  );
}

export default Header;
