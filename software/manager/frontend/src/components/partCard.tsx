import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCamera, faCheckCircle } from "@fortawesome/free-solid-svg-icons"
import { PropertyModel } from "../types"

type Props = {
  id: number
  name: string
  description: string
  imageUrl?: string
  tags: string[]
  properties: PropertyModel[]
  actions: JSX.Element[] | JSX.Element
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

  let image = (
    <div className="w-full h-full bg-base-300 p-16">
      <FontAwesomeIcon icon={faCamera} size="4x" className="text-gray-300 w-full h-full" />
    </div>
  )
    console.log(props.imageUrl)
  if (props.imageUrl) {
    image = <img className="h-full w-full object-contain bg-base-300" src={props.imageUrl} />
  }

  return (
    <div className="card w-full bg-base-100 flex-col sm:flex-row ">
      <figure className="w-full sm:w-64 h-64 bg-base-300 shrink-0">{image}</figure>
      <div className="card-body p-4 sm:p-6">
        <h2 className="card-title">{props.name}</h2>
        <section className="flex gap-1 flex-wrap">
          {renderedTags}
          {renderedProperties}
        </section>
        <p className="whitespace-pre-wrap">{props.description}</p>
        <div className="card-actions justify-end">
          {props.actions}
        </div>
      </div>
    </div>
  )
}
