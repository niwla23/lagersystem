import React from "react"

export default function Boxes() {
  const rowsPerSegment = 4
  const colsPerSegment = 8
  const segmentsX = 3
  const segmentsY = 2

  let boxesRendered = []

  for (let i = 0; i < rowsPerSegment * segmentsX * (colsPerSegment * segmentsY); i++) {
    let isBoxUsed = Math.random() > 0.6
    let isBoxOutOfOrder = Math.random() > 0.9

    let buttonType = "btn-success"
    if (isBoxUsed) buttonType = "btn-warning"
    if (isBoxOutOfOrder) buttonType = "btn-error"

    boxesRendered.push(
      <button className={`btn w-full h-12 col-span-2 ${buttonType}`}>
        <p>{i}</p>
      </button>
    )
    if ((i + 1) % rowsPerSegment === 0) boxesRendered.push(<div className="w-px h-[120%] "></div>)
    if ((i + 1) / colsPerSegment === 12) boxesRendered.push(<div className="w-full col-span-full h-2"></div>)
  }

  return (
    <div className="h-full p-4 w-full overflow-x-clip">
      <p>Click any box number to see details</p>
      <div className="w-screen overflow-x-auto overflow-y-clip h-full">
        <main
          className="grid gap-1 min-w-max h-max"
          style={{ gridTemplateColumns: `repeat(${rowsPerSegment * segmentsX * 2 + segmentsX}, minmax(0, 1fr))` }}
        >
          {boxesRendered}
        </main>
      </div>
    </div>
  )
}
