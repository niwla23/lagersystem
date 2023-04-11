import { useLoaderData, useNavigate } from "react-router-dom"
import { PartModel, UpdatePartData } from "../types"
import { useState } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { AddPartDataForm } from "../components/partAddDataForm"
import * as api from "../api"

export default function EditPart() {
  const navigate = useNavigate()
  const partData = useLoaderData() as PartModel

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

    let updateData: UpdatePartData = {
      name: data.partName,
      description: data.description,
      tags: data.tags.split(","),
      properties: {},
      amount: -1,
      boxId: partData.box?.id,
    }

    // props.submit(creationData, selectedFile)
    api.updatePart(partData.id, updateData, selectedFile).then((r) => {
      console.log("huh")
      if (r.status == 200) {
        navigate("/")
      }
    })
    console.log(updateData)
  }

  const handleDelete = async (e: Event) => {
    e.preventDefault()
    if (!confirm("Really delete?")) {
      return
    }
    let resp = await api.deletePart(partData.id)
    if (resp.status == 200) {
      navigate("/")
    }
  }

  let imgPreview = <span>No image available.</span>
  if (selectedFile) {
    imgPreview = <img src={URL.createObjectURL(selectedFile)} className="h-32 w-32 object-contain" />
  } else if (api.getImageUrl(partData.imageId)) {
    imgPreview = <img src={api.getImageUrl(partData.imageId)} className="h-32 w-32 object-contain" />
  }

  return (
    <section className="flex justify-center flex-grow h-full items-stretch">
      <div className="max-w-7xl w-full p-4">
        <form>
          <div>
            <label className="label">
              <span className="label-text">Part name</span>
            </label>
            <input
              type="text"
              placeholder="Must be unique"
              className="input input-bordered w-full"
              {...register("partName")}
              defaultValue={partData.name}
            />
          </div>
          <div className="form-control">
            <label className="label">
              <span className="label-text">Description</span>
            </label>
            <textarea
              className="textarea textarea-bordered h-24"
              placeholder="Enter the part description here"
              {...register("description")}
              defaultValue={partData.description}
            ></textarea>
          </div>
          <div>
            <label className="label flex-col items-start">
              <div className="label-text py-1">Foto wählen</div>
              <div className="flex w-full gap-2 bg-base-100">
                {imgPreview}
                <input type="file" className="file-input w-full" onChange={handleFileSelected} />
              </div>
            </label>
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
              defaultValue={partData.tags?.map((i) => i.name).join(",")}
            />
          </div>
          <button onClick={handleSubmit(onSubmit)} className="btn btn-primary w-full mt-4">
            Update part
          </button>
          <button onClick={handleDelete} className="btn btn-warning w-full mt-2">
            Delete part
          </button>
        </form>
      </div>
    </section>
  )
}
