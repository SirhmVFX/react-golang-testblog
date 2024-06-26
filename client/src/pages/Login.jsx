import { Link } from "react-router-dom"

function Login() {
  return (
  <>

      <div className="w-3/6 mx-auto flex flex-col gap-8 justify-center p-4">
        <div>
<h1 className="font-bold">Login</h1>
        <p>Please sign in your credentials</p>

        </div>
        <form action="" className="flex flex-col w-full gap-4">
          <div className="flex flex-col w-ful gap-2 ">
            <label htmlFor="username">Username</label>
            <input type="text" placeholder="username" className="p-4 shadow-sm shadow-gray-50 rounded-lg focus:outline-black" /> 
          </div>
 <div className="flex flex-col w-ful gap-2 ">

            <label htmlFor="username">Email</label>
            <input type="email" placeholder="email" className="p-4 shadow-sm shadow-gray-50 rounded-lg focus:outline-black" /> 
          </div>
 <div className="flex flex-col w-ful gap-2 ">
            <label htmlFor="username">Password</label>
            <input type="password" placeholder="password" className="p-4 shadow-sm shadow-gray-50 rounded-lg focus:outline-black" /> 
          </div>
<button className="bg-black text-white p-4 rounded-lg">Sign in</button>
        </form>

        <Link to={"/signup"}>Don&apos;t have an account yet sign up</Link>

      </div>
  </>
  )
}

export default Login
