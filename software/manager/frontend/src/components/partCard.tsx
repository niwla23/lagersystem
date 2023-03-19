import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCheckCircle } from "@fortawesome/free-solid-svg-icons"
import { PropertyModel } from "../types"

type Props = {
  id: number
  name: string
  description: string
  imageUrl: string
  tags: string[]
  properties: PropertyModel[]
  actionText: string
  onActionClick: (id: number) => void
}

export default function PartCard(props: Props) {
  let renderedTags = props.tags.map((tag) => {
    return <span className="badge badge-secondary">{tag}</span>
  })

  let renderedProperties = props.properties.map((property) => {
    return (
      <span className="badge badge-primary pr-0">
        {property.name} <span className="badge badge-accent ml-2">{property.value}</span>
      </span>
    )
  })

  return (
    <div className="card w-full bg-base-100 flex-col sm:flex-row ">
      <figure className="sm:w-1/6 h-64 w-full">
        <img className="h-full w-full object-contain bg-base-300" src={props.imageUrl} />
      </figure>
      <div className="card-body p-4 sm:p-6 sm:p">
        <h2 className="card-title">{props.name}</h2>
        <section className="flex gap-1">
          {renderedTags}
          {renderedProperties}
        </section>
        <p>{props.description}</p>
        <div className="card-actions justify-end">
          <button className="btn btn-primary" onClick={() => props.onActionClick(props.id)}>
            {props.actionText}
          </button>
        </div>
      </div>
    </div>
  )
}
