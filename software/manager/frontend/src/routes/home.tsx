import React, { useEffect } from "react"
import clsx from "clsx"
import PartCard from "../components/partCard"
import * as api from "../api"
import { PartModel } from "../types"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faArrowLeft, faCancel, faDeleteLeft, faPen, faPlug, faTruckRampBox, faXmark } from "@fortawesome/free-solid-svg-icons"
import { useNavigate } from "react-router-dom"
import Swal from "sweetalert2"
import { isAxiosError } from "axios"

export default function Home() {
  const navigate = useNavigate()
  let [parts, setParts] = React.useState<PartModel[]>([])
  let [searchQuery, setSearchQuery] = React.useState("")
  let [filter, setFilter] = React.useState("")
  let [selectedParts, setSelectedParts] = React.useState<number[]>([])

  const loadAllParts = async () => {
    let x = await api.searchParts(searchQuery, filter)
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
      title: "Store box after connect?",
      showDenyButton: true,
      confirmButtonText: "Yes",
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
  }, [searchQuery, filter])

  useEffect(() => {
    let x = setInterval(() => {
      loadAllParts()
    }, 1000)

    return () => {
      clearInterval(x)
    }
  })

  if (!Array.isArray(parts)) {
    parts = []
  }

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
        onTagClick={(tag) => setFilter(`tags: ${tag}`)}
        imageUrl={api.getImageUrl(part.imageId)}
        modifiedAt={new Date(part.updatedAt)}
      />
    )
  })

  return (
    <div className="h-full p-4">
      <div className="input-group w-full">
        <input
          type="text"
          placeholder="Search…"
          className="input input-bordered w-full"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
        />
      </div>
      <div className="input-group w-full pb-4">
        <input
          type="text"
          placeholder="Filter"
          className="input input-bordered w-full"
          value={filter}
          onChange={(e) => setFilter(e.target.value)}
        />

        <div className="dropdown dropdown-left">
          <label tabIndex={0} className="btn">Preset</label>
          <ul tabIndex={0} className="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52">
            <li><a onClick={() => setFilter("hasBox: false")}>Ghost parts</a></li>
            <li><a onClick={() => setFilter("isStored: false && hasBox: true")}>Unstored non-ghost parts</a></li>
            <li><a onClick={() => setFilter("isStored: false")}>Unstored parts</a></li>
            <li><a onClick={() => setFilter("isStored: true")}>Stored parts</a></li>
            <li><a onClick={() => setFilter("")}>No Filter</a></li>
          </ul>
        </div>
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
