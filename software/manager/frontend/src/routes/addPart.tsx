import React, { useState } from "react"
import PartAddDataForm, { AddPartDataForm } from "../components/partAddDataForm"
import ChoosePartAddMode, { AddMode } from "../components/partAddMode"
import PartAddSearch from "../components/partAddSearch"
import { CreatePartData, UpdatePartData } from "../types"
import * as api from "../api"
import { useNavigate } from "react-router-dom"

export default function AddPart() {
  const navigate = useNavigate()
  type AddStage = "selectMode" | "dataForm" | "getting" | "storing" | "confirmPutPartIn"

  const [stage, setStage] = useState<AddStage>("selectMode")
  const [partAddMode, setPartAddMode] = useState<AddMode>("emptyBox")
  const [partCreateData, setPartCreateData] = useState<CreatePartData | null>(null)
  const [partId, setPartId] = useState<number | null>(null)

  const deliverBoxForPutIn = async (boxId: string) => {
    setStage("getting")
    await api.deliverBox(boxId)
    setStage("confirmPutPartIn")
  }

  const storeBox = async () => {
    setStage("storing")
    const resp = await api.clearScanner()
    if (!partId) {
      alert("partId not set")
      return
    }
    await api.updatePart(partId, { ...(partCreateData as UpdatePartData), boxId: resp.boxId }, undefined)
    alert("successfully stored box")
    navigate("/")
  }

  const submitForm = async (form: CreatePartData, image: File | undefined) => {
    // save form data for later and create part in the database. Do not yet assign a box
    setPartCreateData(form)
    try {
      const resp = await api.createPart(form, image)
      setPartId(resp.id)
    } catch (e: any) {
      alert(e.response.data.message)
      return
    }

    switch (partAddMode) {
      case "emptyBox":
        // user wants to store part in an empty box, find one and deliver it
        try {
          const emptyBox = await api.getEmptyBox()
          deliverBoxForPutIn(emptyBox.id)
        } catch (e) {
          alert("no free box found")
          navigate("/")
          return
        }
        break

      case "createOnly":
        // we are done here
        navigate("/")
        break

      case "storeOnly":
        // user already has the box in their hands and just wants to store the box and assign it to the part
        setStage("confirmPutPartIn")
        break
    }
  }

  let content = <div></div>
  if (stage === "selectMode") {
    content = <ChoosePartAddMode modeChosen={partAddMode} setModeChosen={setPartAddMode} submit={() => setStage("dataForm")} />
  } else if (stage === "dataForm") {
    content = <PartAddDataForm submit={submitForm} />
  } else if (stage === "getting") {
    content = (
      <div className="h-full grid place-items-center">
        <div className="w-full">
          <h2 className="font-bold text-xl">Getting box...</h2>
          <progress className="progress w-full"></progress>
        </div>
      </div>
    )
  } else if (stage === "storing") {
    content = (
      <div className="h-full grid place-items-center">
        <div className="w-full">
          <h2 className="font-bold text-xl">Storing part "{partCreateData?.name}"...</h2>
          <progress className="progress w-full"></progress>
        </div>
      </div>
    )
  } else if (stage === "confirmPutPartIn") {
    content = (
      <div className="h-full grid place-items-center">
        <div className="w-full flex flex-col gap-2 justify-center">
          <h2 className="font-bold text-xl text-center">Please put the parts in the box and place it on the scanner</h2>
          <button className="btn btn-warning whitespace-nowrap" onClick={() => storeBox()}>
            Store box now
          </button>
        </div>
      </div>
    )
  }

  return (
    <section className="flex justify-center flex-grow h-full items-stretch">
      <div className="max-w-7xl w-full p-4">{content}</div>
    </section>
  )
}
