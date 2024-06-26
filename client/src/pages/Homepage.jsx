import Blogpost from "../components/Blogpost"

function Homepage() {
  return (
  <>

      <div className="flex flex-col items-center p-8  ">

        <h1>
          All Post

                 </h1>
 <div className="grid-cols-2 w-5/6 mx-auto p-4 grid gap-4">
            <Blogpost />
            <Blogpost />
            <Blogpost />
            <Blogpost />
            <Blogpost />
            <ilogpost />
          </div>

      </div>
  </>
  )
}

export default Homepage
