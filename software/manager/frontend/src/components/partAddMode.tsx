// this component allows the user to choose how to store the newly added part.
export type AddMode = "createOnly" | "storeOnly"

let ModeCard = (props: { title: string; text: string; active: boolean; onClick: () => void }) => {
  return (
    <div
      className={`card w-full bg-base-100 shadow-xl transition-all duration-200 ring-primary ${props.active ? "ring-4" : ""}`}
      onClick={props.onClick}
    >
      <div className="card-body">
        <h2 className="card-title">{props.title}</h2>
        <p>{props.text}</p>
      </div>
    </div>
  )
}

type Props = {
  modeChosen: AddMode
  setModeChosen: (modeChosen: AddMode) => void
  submit: () => void
}

export default function PartAddMode(props: Props) {
  return (
    <div className="flex flex-col w-full gap-4">
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <ModeCard
          title="Only create part"
          text="Creates the part in the database but does not yet store it in the warehouse"
          active={props.modeChosen === "createOnly"}
          onClick={() => props.setModeChosen("createOnly")}
        />
        <ModeCard
          title="Add & Store"
          text="Create the part, link it to the box that is placed on the scanner, then store that box at a free position"
          active={props.modeChosen === "storeOnly"}
          onClick={() => props.setModeChosen("storeOnly")}
        />
      </div>
      <button className="btn w-full" onClick={props.submit}>
        Next
      </button>
    </div>
  )
}
