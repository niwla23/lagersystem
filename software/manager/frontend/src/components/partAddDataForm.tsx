import { useRef, useState } from "react"
import { useForm, SubmitHandler } from "react-hook-form"
import { CreatePartData } from "../types"

export type AddPartDataForm = {
  partName: string
  description: string
  tags: string
  box: number
  section: number
}

type Props = {
  submit: (form: CreatePartData, image: File | undefined) => void
}

export default function PartAddDataForm(props: Props) {
  let [selectedFile, setSelectedFile] = useState<File | undefined>(undefined)
  const handleFileSelected = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setSelectedFile(e.target.files?.item(0) || undefined)
  }

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AddPartDataForm>()

  const onSubmit: SubmitHandler<AddPartDataForm> = (data) => {
    console.log(data)

    let creationData: CreatePartData = {
      name: data.partName,
      description: data.description,
      tags: data.tags.split(","),
      properties: {},
      amount: -1,
    }

    props.submit(creationData, selectedFile)
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
          <span className="label-text">Foto wählen</span>
        </label>
        <input type="file" className="file-input w-full" onChange={handleFileSelected} />
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

      <button onClick={handleSubmit(onSubmit)} className="btn btn-primary w-full mt-4">
        Store part
      </button>
    </form>
  )
}
