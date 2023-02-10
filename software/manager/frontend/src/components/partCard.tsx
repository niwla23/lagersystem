import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons"

type Props = {
  id: number
  name: string
  actionText: string
  onActionClick: (id: number) => void
}

export default function PartCard(props: Props) {
  return (
    <div className="card w-full bg-base-100 flex-row p-6">
      <figure className="">
        <img
          className="h-32 w-full object-cover"
          src={`https://cdn-reichelt.de/resize/150/150/web/artikel_ws/EB00/CREALITY3D_LD002R_01.jpg?20210317`}
          alt="Shoes"
        />
      </figure>
      <div className="card-body p-0">
        <h2 className="card-title">{props.name}</h2>
        <section className="flex gap-1">
          <span className="badge">Tech</span>
          <span className="badge">Resistor</span>
          <span className="badge">5% Tolerance</span>
        </section>
        <p>If a dog chews shoes whose shoes does he choose?</p>
        <div className="card-actions justify-end">
          <button className="btn btn-primary" onClick={() => props.onActionClick(props.id)}>
            {props.actionText}
          </button>
        </div>
      </div>
    </div>
  )
}
