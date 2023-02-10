import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faCheckCircle, faSearch } from "@fortawesome/free-solid-svg-icons"
import PartCard from "./partCard"

type Props = {
  submit: () => void
}

export default function PartAddSearch(props: Props) {
  return (
    <section className="w-full">
      <div className="input-group pb-4">
        <input type="text" placeholder="Search…" className="input input-bordered w-full" />
        <button className="btn btn-square">
          <FontAwesomeIcon icon={faSearch} />
        </button>
      </div>
      <main className="flex flex-col gap-2">
        <PartCard id={1} onActionClick={props.submit} actionText="Choose" name="Widerstand" />
        <PartCard id={2} onActionClick={props.submit} actionText="Choose" name="M4 Schraube" />
      </main>

      <button onClick={props.submit} className="btn btn-primary w-full mt-4">
        Store part
      </button>
    </section>
  )
}
