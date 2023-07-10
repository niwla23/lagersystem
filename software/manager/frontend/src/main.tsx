import React from "react"
import ReactDOM from "react-dom/client"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import "./index.css"
import Layout from "./layout"
import AboutRoute from "./routes/about"
import AddPart from "./routes/addPart"
import Boxes from "./routes/boxes"
import HomeRoute from "./routes/home"
import EditPart from "./routes/editPart"
import * as api from "./api" 
import StoreBox from "./routes/storeBox"

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: "/",
        element: <HomeRoute />,
      },
      {
        path: "/boxes",
        element: <Boxes />,
        loader: async ({params}) => {
          return await api.getPositions();
        },
      },
      {
        path: "/boxes/store",
        element: <StoreBox />,
      },
      {
        path: "/parts/add",
        element: <AddPart />,
      },
      {
        path: "/parts/edit/:partId",
        element: <EditPart />,
        loader: async ({params}) => {
          return await api.getPartById(Number(params.partId));
        },
      },
      {
        path: "/about",
        element: <AboutRoute />,
      },
    ],
  },
])

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
