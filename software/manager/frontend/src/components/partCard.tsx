import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faBox, faCamera, faCheckCircle, faGhost, faStore, faWarehouse } from "@fortawesome/free-solid-svg-icons"
import { PropertyModel } from "../types"

type Props = {
  id: number
  name: string
  description: string
  imageUrl?: string
  tags: string[]
  properties: PropertyModel[]
  actions: JSX.Element[] | JSX.Element
  boxId?: string
  positionId?: number
  checked: boolean
  modifiedAt: Date
  onCheckChanged: (checked: boolean) => void
  onTagClick: (tagName: string) => void
}

export default function PartCard(props: Props) {
  let renderedTags = props.tags.map((tag) => {
    return (
      <span className="badge badge-secondary" key={tag} onClick={() => props.onTagClick(tag)}>
        {tag}
      </span>
    )
  })

  let renderedProperties = props.properties.map((property) => {
    return (
      <span className="badge badge-primary pr-0" key={property.id}>
        {property.name} <span className="badge badge-accent ml-2">{property.value}</span>
      </span>
    )
  })

  let image = (
    <div className="w-full h-full bg-base-300 p-16">
      <FontAwesomeIcon icon={faCamera} size="4x" className="text-gray-300 w-full h-full" />
    </div>
  )
  if (props.imageUrl) {
    image = <img className="h-full w-full object-cover bg-base-300" loading="lazy" src={props.imageUrl} />
  }

  let icon = faGhost
  let tooltip = "No box assigned"
  if (props.positionId) {
    icon = faCheckCircle
    tooltip = "Ready for delivery"
  } else if (props.boxId) {
    tooltip = "Not stored"
    icon = faBox
  }

  return (
    <div className="card w-full bg-base-100 flex-col sm:flex-row ">
      <figure className="w-full sm:w-64 h-64 bg-base-300 shrink-0">{image}</figure>
      <div className="card-body p-4 sm:p-6">
        <h2 className="card-title">
          <input
            type="checkbox"
            className="checkbox checkbox-xs"
            checked={props.checked}
            onChange={(e) => props.onCheckChanged(e.target.checked)}
          />
          {props.name}
        </h2>
        <section className="flex gap-1 flex-wrap">
          <div className="tooltip pr-1 transform -translate-y-0.5" data-tip={tooltip}>
            <FontAwesomeIcon icon={icon} className="text-gray-500" />
          </div>
          {new Date().getTime() - props.modifiedAt.getTime() < 20 * 60 * 1000 && <span className="badge badge-info font-bold">NEW</span>}
          {renderedTags}
          {renderedProperties}
        </section>
        <p className="whitespace-pre-wrap">{props.description}</p>
        <div className="card-actions justify-end">{props.actions}</div>
      </div>
    </div>
  )
}
