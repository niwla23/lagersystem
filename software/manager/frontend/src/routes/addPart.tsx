import PartAddDataForm from "../components/partAddDataForm"
import { CreatePartData } from "../types"
import * as api from "../api"
import { useNavigate } from "react-router-dom"
import Swal from "sweetalert2"

export default function AddPart() {
  const navigate = useNavigate()

  const connectToBox = async (partId: number, storeAfterwards: boolean) => {
    await api.bulkLinkParts([partId], storeAfterwards)
    navigate("/")
  }

  const submitForm = async (form: CreatePartData, image: File | undefined) => {
    try {
      const resp = await api.createPart(form, image)

      let result = await Swal.fire({
        title: "Store Box Now?",
        icon: "question",
        showConfirmButton: true,
        showDenyButton: true,
        confirmButtonText: "Yes",
        denyButtonText: "No"
      })

      await connectToBox(resp.id, result.isConfirmed)
      navigate("/")
    } catch (e: any) {
      alert(e.response.data.message)
      return
    }
  }

  return (
    <section className="flex justify-center flex-grow h-full items-stretch">
      <div className="max-w-7xl w-full p-4">
        <PartAddDataForm submit={submitForm} />
      </div>
    </section>
  )
}
