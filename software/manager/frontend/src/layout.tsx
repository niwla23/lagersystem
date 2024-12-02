import React from "react"
import { Link, Outlet, useLocation } from "react-router-dom"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faChevronDown, faAngleRight, faBars, faAdd, faBell } from "@fortawesome/free-solid-svg-icons"

export default function Layout() {
  let location = useLocation()

  // this component is used to highlight the current page in the menu
  let LinkWithStyle = (props: { to: string; children: JSX.Element | JSX.Element[] | string }) => {
    let style = location.pathname === props.to ? "bg-secondary text-secondary-content" : ""

    return (
      <Link to={props.to} className={style}>
        {props.children}
      </Link>
    )
  }

  let menu = (position: "vertical" | "horizontal") => {
    let expandIcon = position === "vertical" ? faAngleRight : faChevronDown

    return (
      <>
        <li>
          <LinkWithStyle to={"/"}>Parts</LinkWithStyle>
        </li>
        <li>
          <LinkWithStyle to={"/boxes"}>Boxes</LinkWithStyle>
        </li>
        <li>
          <LinkWithStyle to={"/boxes/store"}>Store box</LinkWithStyle>
        </li>
        {/* <li>
          <a>
            Parent
            <FontAwesomeIcon icon={expandIcon} />
          </a>
          <ul className="p-2 bg-base-100">
            <li>
              <a>Submenu 1</a>
            </li>
            <li>
              <a>Submenu 2</a>
            </li>
          </ul>
        </li> */}
      </>
    )
  }

  return (
    <div className="bg-base-200 min-h-screen flex flex-col">
      <div className="navbar bg-base-100">
        <section className="navbar-start w-full">
          {/* mobile menu */}
          <div className="dropdown">
            <label tabIndex={0} className="btn btn-ghost lg:hidden">
              <FontAwesomeIcon icon={faBars} size="lg" />
            </label>
            <ul tabIndex={0} className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
              {menu("vertical")}
            </ul>
          </div>
          <a className="btn btn-ghost normal-case text-xl">Lagersystem</a>
          {/* desktop menu */}
          <ul className="menu menu-horizontal px-1 hidden lg:flex z-10">{menu("horizontal")}</ul>
        </section>
        <section className="navbar-end gap-2">
          <Link to={"/parts/add"} className="btn btn-ghost">
            <FontAwesomeIcon icon={faAdd} />
          </Link>
          <button className="btn btn-ghost">
            <FontAwesomeIcon icon={faBell} />
          </button>
          <div className="avatar pl-2">
            <div className="h-8 rounded-full">
              <img src={`https://www.gravatar.com/avatar/43661af400162079e6dcadef7debd43a`} />
            </div>
          </div>
        </section>
      </div>
      <Outlet />
    </div>
  )
}
