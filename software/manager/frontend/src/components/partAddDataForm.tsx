import { useForm, SubmitHandler } from "react-hook-form"

export type AddPartDataForm = {
  partName: string
  description: string
  tags: string
  box: number
  section: number
}

type Props = {
  submit: (form: AddPartDataForm) => void
}

export default function PartAddDataForm(props: Props) {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AddPartDataForm>()
  const onSubmit: SubmitHandler<AddPartDataForm> = (data) => {
    console.log(data)
    props.submit(data)
  }

  return (
    <form>
      <div>
        <label className="label">
          <span className="label-text">Part name</span>
        </label>
        <input type="text" placeholder="Must be unique" className="input input-bordered w-full" {...register("partName")} />
      </div>
      <div className="form-control">
        <label className="label">
          <span className="label-text">Description</span>
        </label>
        <textarea
          className="textarea textarea-bordered h-24"
          placeholder="Enter the part description here"
          {...register("description")}
        ></textarea>
      </div>
      <div>
        <label className="label">
          <span className="label-text">Fotos wählen</span>
        </label>
        <input type="file" className="file-input w-full" multiple />
      </div>
      <div>
        <label className="label">
          <span className="label-text">Tags</span>
        </label>
        <input
          type="text"
          placeholder="Seperated by comma (e.g.: resistor,electronics)"
          className="input input-bordered w-full"
          {...register("tags")}
        />
      </div>
      <div>
        <label className="label">
          <span className="label-text">Box</span>
        </label>
        <input type="text" placeholder="Must be unique" className="input input-bordered w-full" {...register("box")} />
      </div>
      <div className="form-control w-full">
        <label className="label">
          <span className="label-text">Section</span>
        </label>
        <select className="select select-bordered" {...register("section")}>
          <option disabled>Pick one</option>
          <option>Section 1</option>
          <option>Section 2</option>
          <option>Section 3</option>
        </select>
      </div>

      <button onClick={handleSubmit(onSubmit)} className="btn btn-primary w-full mt-4">
        Store part
      </button>
    </form>
  )
}
