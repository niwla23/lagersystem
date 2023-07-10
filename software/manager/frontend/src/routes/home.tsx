import React, { useEffect } from "react"
import clsx from "clsx"
import PartCard from "../components/partCard"
import * as api from "../api"
import { PartModel } from "../types"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faArrowLeft, faCancel, faPen, faPlug, faTruckRampBox, faXmark } from "@fortawesome/free-solid-svg-icons"
import { useNavigate } from "react-router-dom"
import Swal from "sweetalert2"
import { isAxiosError } from "axios"

export default function Home() {
  const navigate = useNavigate()
  let [parts, setParts] = React.useState<PartModel[]>([])
  let [searchQuery, setSearchQuery] = React.useState("")
  let [selectedParts, setSelectedParts] = React.useState<number[]>([])

  const loadAllParts = async () => {
    let x = await api.searchParts(searchQuery, "")
    setParts(x)
  }

  const handleSelectChanged = (partId: number, selected: boolean) => {
    if (selected) {
      setSelectedParts([...selectedParts, partId])
    } else {
      setSelectedParts(selectedParts.filter((x) => x != partId))
    }
  }

  const deliverPart = async (partId: number) => {
    try {
      await api.deliverPart(partId)
      loadAllParts()
    } catch (e) {
      if (isAxiosError(e)) {
        Swal.fire({
          title: "Error",
          text: e.response?.data?.message || e.message,
          icon: "error",
        })
      }
    }
  }

  const deliverSelectedParts = async () => {
    let boxesWanted = selectedParts.map((partId) => parts.find((x) => x.id == partId)?.box?.id)
    let boxesWantedUnique = [...new Set(boxesWanted)]
    let availableSlots = await api.getFreeIoSlots()

    if (boxesWantedUnique.length > availableSlots.length) {
      Swal.fire({ title: "Error", text: "Not enought IO slots to deliver these boxes", icon: "error" })
      return
    }
    selectedParts.forEach(async (partId) => {
      await deliverPart(partId)
    })
  }

  const connectSelectedPartsToBox = async () => {
    let result = await Swal.fire({
      title: 'Store box after connect?',
      showDenyButton: true,
      confirmButtonText: 'Yes',
      denyButtonText: `No`,
    })
    let storeAfterwards = result.isConfirmed

    const swalWithLoader = Swal.mixin({
      title: "Linking parts to box...",
      allowOutsideClick: false,
      customClass: {},
      didOpen: () => {
        Swal.showLoading()
      },
    })

    swalWithLoader.fire()

    try {
      await api.bulkLinkParts(selectedParts, storeAfterwards)
      loadAllParts()
      swalWithLoader.close()
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

  useEffect(() => {
    loadAllParts()
  }, [searchQuery])

  useEffect(()=>{
    let x = setInterval(()=>{
      loadAllParts()
    }, 1000)

    return () => {
      clearInterval(x)
    }
  })

  let renderedParts = parts.map((part) => {
    let tags = part.tags ? part.tags.map((tag) => tag.name) : []
    return (
      <PartCard
        id={part.id}
        actions={[
          <button
            key={"deliver"}
            className={clsx("btn btn-primary flex-grow sm:flex-grow-0", { "btn-disabled": !part.box || !part.box.position })}
            onClick={() => deliverPart(part.id)}
          >
            Deliver
          </button>,
          <button
            key={"edit"}
            className="btn btn-secondary"
            onClick={() => {
              navigate(`/parts/edit/${part.id}`)
            }}
          >
            <FontAwesomeIcon icon={faPen} />
          </button>,
        ]}
        name={part.name}
        key={String(part.id)}
        description={part.description || "No description"}
        tags={tags}
        properties={part.properties || []}
        boxId={part.box?.id}
        positionId={part.box?.position?.id}
        checked={selectedParts.includes(part.id)}
        onCheckChanged={(selected) => handleSelectChanged(part.id, selected)}
        imageUrl={api.getImageUrl(part.imageId)}
        modifiedAt={new Date(part.updatedAt)}
      />
    )
  })

  return (
    <div className="h-full p-4">
      <div className="input-group w-full pb-4">
        <input
          type="text"
          placeholder="Search…"
          className="input input-bordered w-full"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
        />
        <button className="btn btn-primary btn-square">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </button>
      </div>
      <main className="grid gap-2">{renderedParts}</main>
      <div
        className={clsx("fixed right-0 top-0 h-full flex flex-col justify-center transform transition-all", {
          "translate-x-40": selectedParts.length == 0,
        })}
      >
        <div className="bg-accent text-black p-4 flex flex-col gap-4 text-center items-center">
          <p className="font-bold">{selectedParts.length}</p>
          <button className="tooltip tooltip-left" data-tip="Connect box" onClick={connectSelectedPartsToBox}>
            <FontAwesomeIcon icon={faPlug} />
          </button>
          <button className="tooltip tooltip-left" data-tip="Deliver" onClick={deliverSelectedParts}>
            <FontAwesomeIcon icon={faArrowLeft} />
          </button>
          <button className="tooltip tooltip-left" data-tip="Cancel Selection" onClick={() => setSelectedParts([])}>
            <FontAwesomeIcon icon={faXmark} />
          </button>
        </div>
      </div>
    </div>
  )
}
