import Layout from '@/layout'
import { Navigate, type RouteObject } from 'react-router-dom'

const routes: RouteObject[] = [
  // 实现重定向
  {
    path: "/",
    element: <Navigate to="/page-a" replace={true}/>
  },
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: "page-a",
        async lazy() {
          const com = await import("@/page/page-a")
          return { Component: com.default }
        }
      },
      {
        path: "page-b",
        async lazy() {
          const com = await import("@/page/page-b")
          return { Component: com.default }
        }
      }
    ]
  }

]

export default routes
