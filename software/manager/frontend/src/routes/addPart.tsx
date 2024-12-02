import PartAddDataForm from "../components/partAddDataForm"
import { CreatePartData } from "../types"
import * as api from "../api"
import { useNavigate } from "react-router-dom"

export default function AddPart() {
  const navigate = useNavigate()

  const submitForm = async (form: CreatePartData, image: File | undefined) => {
    try {
      await api.createPart(form, image)
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
