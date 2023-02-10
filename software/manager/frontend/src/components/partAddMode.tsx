import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons"

export type AddMode = "freeSection" | "storeWithPart" | "storeOnly"

let ModeCard = (props: { title: string; text: string; active: boolean; onClick: () => void }) => {
  return (
    <div
      className={`card w-full bg-base-100 shadow-xl transition-all duration-200 ${props.active ? "bg-primary" : ""}`}
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
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <ModeCard
          title="Give me a free section"
          text="The system will find an unused section and deliver the box for you to put in the part(s)"
          active={props.modeChosen === "freeSection"}
          onClick={() => props.setModeChosen("freeSection")}
        />
        <ModeCard
          title="I want to put this in the same section as another part"
          text="You will be able to choose the part this should be stored together with, then the box is delivered for you to put in the part(s)"
          active={props.modeChosen === "storeWithPart"}
          onClick={() => props.setModeChosen("storeWithPart")}
        />
        <ModeCard
          title="I have already taken out the box in which I want to put it"
          text="The Code of the Box you put in the input tray will be scanned and stored where it belongs"
          active={props.modeChosen === "storeOnly"}
          onClick={() => props.setModeChosen("storeOnly")}
        />
      </div>
      <button className="btn w-full" onClick={props.submit}>Next</button>
    </div>
  )
}
