import { Outlet, Link } from 'react-router-dom'

const Layout: React.FC = () => {


  return (
    <div>
      <Link to="/page-a">page a</Link>
      <Link to="/page-b">page b</Link>
      <Outlet />
    </div>
  )
}

export default Layout
