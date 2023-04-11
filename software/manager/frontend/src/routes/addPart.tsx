import React, { useState } from "react"
import PartAddDataForm, {AddPartDataForm} from "../components/partAddDataForm"
import ChoosePartAddMode, { AddMode } from "../components/partAddMode"
import PartAddSearch from "../components/partAddSearch"
import { CreatePartData } from "../types"
import * as api from "../api"


export default function AddPart() {
  type AddStage = "selectMode" | "dataForm" | "searchPart" | "getting" | "storing" | "confirmPutPartIn"

  const [stage, setStage] = useState<AddStage>("selectMode")
  const [partAddMode, setPartAddMode] = useState<AddMode>("freeSection")
  const [partCreateData, setPartCreateData] = useState<CreatePartData | null>(null)


  let content = <div></div>
  if (stage === "selectMode") {
    content = <ChoosePartAddMode modeChosen={partAddMode} setModeChosen={setPartAddMode} submit={() => setStage("dataForm")} />
  } else if (stage === "searchPart") {
    content = content = <PartAddSearch submit={() => setStage("getting")} />
  } else if (stage === "dataForm") {
    content = (
      <PartAddDataForm
        submit={async (form, image) => {
          setPartCreateData(form)
          try {
            let e2 = await api.createPart(form, image)
          } catch(e: any) {
            alert(e.response.data.message)
            return
          }
          switch (partAddMode) {
            case "freeSection":
              setStage("getting")
              break

            case "storeWithPart":
              setStage("searchPart")
              break

            case "storeOnly":
              setStage("confirmPutPartIn")
              break
          }
        }}
      />
    )
  } else if (stage === "getting") {
    content = (
      <div className="h-full grid place-items-center">
        <div className="w-full">
          <h2 className="font-bold text-xl">Getting box...</h2>
          <progress className="progress w-full"></progress>
          <button className="btn btn-error" onClick={() => setStage("confirmPutPartIn")}>
            Done (to be automated)
          </button>
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
        <div className="w-min flex flex-col gap-2 justify-center">
          <h2 className="font-bold text-xl text-center whitespace-nowrap">Please put the parts in the box</h2>
          <button className="btn btn-warning whitespace-nowrap" onClick={() => setStage("storing")}>
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
