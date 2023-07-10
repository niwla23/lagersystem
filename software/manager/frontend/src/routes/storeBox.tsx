import { useState } from "react"
import * as api from "../api"

export default function StoreBox() {
  const [stage, setStage] = useState("start")

  const storeBox = async () => {
    setStage("storing")
    const resp = await api.clearIo()
    setStage("done")
  }

  let mainContent = <></>
  switch (stage) {
    case "start":
      mainContent = (
        <button className="btn btn-primary w-full" onClick={storeBox}>
          Store box
        </button>
      )
      break

    case "storing":
      mainContent = <progress className="progress w-full"></progress>
      break

    case "done":
      mainContent = (
        <>
          <p>Box stored.</p>
          <button className="btn btn-primary w-full" onClick={storeBox}>
            Store another box
          </button>
        </>
      )
      break
  }

  return (
    <section className="flex-grow h-full grid items-center justify-center">
      <div className="max-w-2xl">
        <h1 className="text-xl font-bold">Clearing IO Slots...</h1>
        {mainContent}
      </div>
    </section>
  )
}
