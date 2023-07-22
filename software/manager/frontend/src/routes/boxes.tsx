import React, { useEffect } from "react"
import { Link, useLoaderData } from "react-router-dom"
import { OperatorPositionsResponse } from "../types"
import Swal from "sweetalert2"
import withReactContent from "sweetalert2-react-content"
import * as api from "../api"

function makeCanvas(positions: OperatorPositionsResponse) {
  const canvas = document.getElementById("board") as HTMLCanvasElement | null
  if (!canvas) return
  const ctx = canvas.getContext("2d")
  if (!ctx) return
  const width = 1100
  const height = 800

  ctx.fillStyle = "rgb(50, 50, 50)"
  ctx.fillRect(0, 0, canvas.width, canvas.height)
  for (const [posId, coord] of Object.entries(positions.positions)) {
    const x = width - coord["x"]
    const y = height - coord["y"]

    ctx.fillStyle = coord.box ? (coord.box.parts ? "rgb(200, 100, 0)" : "rgb(0, 100, 50)") : "rgb(0, 0, 0)"
    ctx.fillRect(x, y, 68, 36)
    ctx.fillStyle = "rgb(255, 255, 255)"
    ctx.fillText(posId, x + 20, y + 20)
  }
}

export default function Boxes() {
  const positions = useLoaderData() as OperatorPositionsResponse
  const MySwal = withReactContent(Swal)
  const [selectedPosition, setSelectedBox] = React.useState<string | undefined>(undefined)

  const canvasRef = React.useRef<HTMLCanvasElement>(null)

  const deliverBox = async (boxId: string) => {
    await api.deliverBox(positions.positions[boxId].box!.id)
  }

  const showPosition = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!selectedPosition) return
    let boxId = positions.positions[selectedPosition].box?.id
    MySwal.fire({
      title: `Position ${selectedPosition}`,
      html: (
        <div className="w-full text-left text-white">
          <p className="mb-2">
            Box: <small className="bg-base-100 text-xs">{boxId || "None"}</small>
          </p>
          <div className="w-full flex gap-2">
            <button className="btn btn-error flex-1" disabled={!boxId}>
              Box is not stored here
            </button>
            <button className="btn btn-primary flex-1" onClick={() => deliverBox(selectedPosition)} disabled={!boxId}>
              Deliver box
            </button>
          </div>
        </div>
      ),
    })
  }

  const deliverEmpty = async (amount: number) => {
    for (let i = 0; i < amount; i++) {
      let box = await api.getEmptyBox()
      await api.deliverBox(box.id)
    }
  }

  useEffect(() => {
    makeCanvas(positions)
  }, [canvasRef])

  return (
    <div className="w-full overflow-x-clip h-screen overflow-y-clip md:h-full">
      <div className="bg-base-100 p-4 m-4">
        <p>Enter any position id to see details</p>
        <div className="flex gap-2">
          <form className="input-group" onSubmit={showPosition}>
            <input
              type="number"
              placeholder="69"
              className="input input-bordered w-full"
              value={selectedPosition}
              onChange={(e) => setSelectedBox(e.target.value)}
            />
            <button className="btn" type="submit">
              Go
            </button>
          </form>
          <Link to="/boxes/scan">
            <button className="btn">
              Scanner
            </button>
          </Link>
        </div>
      </div>
      <div className="bg-base-100 p-4 m-4">
        <p>Deliver empty boxes</p>
        <div className="flex gap-2 max-w-lg">
          <button className="btn flex-1" onClick={() => deliverEmpty(1)}>
            Deliver 1
          </button>
          <button className="btn flex-1" onClick={() => deliverEmpty(2)}>
            Deliver 2
          </button>
          <button className="btn flex-1" onClick={() => deliverEmpty(3)}>
            Deliver 3
          </button>
        </div>
      </div>
      <div className="w-screen overflow-auto h-screen pb-64 md:h-auto md:overflow-x-clip md:m-4">
        <canvas width="1200" height="900" id="board" ref={canvasRef} />
      </div>
    </div>
  )
}
