import { useLoaderData, useNavigate, useRevalidator } from "react-router-dom"
import { PartModel, UpdatePartData } from "../types"
import { MouseEventHandler, useState } from "react"
import { Controller, SubmitHandler, useForm } from "react-hook-form"
import { AddPartDataForm } from "../components/partAddDataForm"
import * as api from "../api"
import TagSelector from "../components/tagSelector"
import Swal from "sweetalert2"
import { isAxiosError } from "axios"

export default function EditPart() {
  const navigate = useNavigate()
  let revalidator = useRevalidator()

  const partData = useLoaderData() as PartModel

  let [selectedFile, setSelectedFile] = useState<File | undefined>(undefined)
  const handleFileSelected = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setSelectedFile(e.target.files?.item(0) || undefined)
  }

  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<AddPartDataForm>()

  const onSubmit: SubmitHandler<AddPartDataForm> = (data) => {
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
      if (r.status == 200) {
        navigate("/")
      }
    })
  }

  const handleDelete = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault()
    if (!confirm("Really delete?")) {
      return
    }
    let resp = await api.deletePart(partData.id)
    if (resp.status == 200) {
      navigate("/")
    }
  }

  const deliverEmptyBox = async () => {
    try {
      const emptyBox = await api.getEmptyBox()
      await api.deliverBox(emptyBox.id)
    } catch (e) {
      alert("no free box found")
      return
    }
  }

  const connectToBoxInScanner = async () => {
    // Show loader before making the API call
    const swalWithLoader = Swal.mixin({
      title: "Scanning box...",
      allowOutsideClick: false,
      customClass: {},
      didOpen: () => {
        Swal.showLoading()
      },
    })

    swalWithLoader.fire()

    try {
      const resp = await api.getBoxFromScanner()
      await api.updatePart(
        partData.id,
        {
          name: partData.name,
          description: partData.description,
          properties: {},
          amount: partData.amount,
          boxId: resp.boxId,
          tags: partData.tags?.map((i) => i.name) || [],
        },
        undefined
      )

      // Hide the loader after API calls are complete
      swalWithLoader.close()
      revalidator.revalidate()
    } catch (error) {
      if (isAxiosError(error)) {
        Swal.fire({
          icon: "error",
          title: "Oops...",
          text: error.response?.data?.message || error.message,
        })
      } else {
        Swal.fire("error")
      }
    }
  }

  const disconnectBox = async () => {
    await api.updatePart(
      partData.id,
      {
        name: partData.name,
        description: partData.description,
        properties: {},
        amount: partData.amount,
        boxId: null,
        tags: partData.tags?.map((i) => i.name) || [],
      },
      undefined
    )
    Swal.fire({
      icon: "success",
      title: "Box disconnected",
    })
    revalidator.revalidate()
  }

  let imgPreview = <span>No image available.</span>
  if (selectedFile) {
    imgPreview = <img src={URL.createObjectURL(selectedFile)} className="h-32 w-32 object-contain" />
  } else if (api.getImageUrl(partData.imageId)) {
    imgPreview = <img src={api.getImageUrl(partData.imageId)} className="h-32 w-32 object-contain" />
  }

  let boxManagementState = "No box assigned"
  if (partData.box?.position) {
    boxManagementState = `Box ${partData.box.id} at position ${partData.box.position}`
  } else if (partData.box?.id) {
    boxManagementState = "Box " + partData.box.id
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
            {/* <input
              type="text"
              placeholder="Seperated by comma (e.g.: resistor,electronics)"
              className="input input-bordered w-full"
              {...register("tags")}
              defaultValue={partData.tags?.map((i) => i.name).join(",")}
            /> */}
            <Controller
              control={control}
              rules={{ required: true, minLength: 1 }}
              name="tags"
              render={({ field }) => <TagSelector label="Tags" {...field} error={errors.tags?.type?.toString()} />}
              defaultValue={partData.tags?.map((i) => i.name).join(",")}
            />
          </div>
          <div className="flex gap-2 mt-2">
            <button onClick={handleDelete} className="btn btn-warning flex-1">
              Delete part
            </button>
            <button onClick={handleSubmit(onSubmit)} className="btn btn-primary flex-1">
              Update part
            </button>
          </div>
        </form>
        <div className="mt-16">
          <h2 className="font-bold text-lg pb-2">Box management</h2>
          <p>{boxManagementState} </p>
          <div className="flex flex-col gap-2">
            <button onClick={deliverEmptyBox} className="btn btn-primary w-full">
              Deliver empty box
            </button>
            <button onClick={connectToBoxInScanner} className="btn btn-primary w-full">
              Part is in the scanner, save it
            </button>
            <button onClick={disconnectBox} className="btn btn-error w-full">
              Part is not in this box
            </button>
          </div>
        </div>
      </div>
    </section>
  )
}
