
import { faPen } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React, { useState } from 'react';
import { OnResultFunction, QrReader } from 'react-qr-reader';
import { useNavigate } from 'react-router-dom';
import Swal from 'sweetalert2';
import * as api from "../api"
import PartCard from '../components/partCard';
import { BoxModel } from '../types';

// react state wasnt nice
declare global {
  interface Window { scanningBlocked: boolean; }
}

export const ScanBox = () => {
  const navigate = useNavigate()
  const [scannedId, setScannedId] = useState('No result');
  const [boxData, setBoxData] = useState<BoxModel | null>(null)

  const loadBoxData = async (boxId: string) => {
    try {
      const boxData = await api.getBoxById(boxId)
      setBoxData(boxData)
    } catch (e) {
      setBoxData(null)
    }
  }

  const onResult: OnResultFunction = (result, error) => {
    if (!!result && !window.scanningBlocked) {
      window.scanningBlocked = true
      setTimeout(() => window.scanningBlocked = false, 1000)
      Swal.fire({ title: "Scanned", icon: "success", timer: 600 })
      loadBoxData(result.getText())
      setScannedId(result.getText())
    }

    if (!!error) {
      console.info(error);
    }
  }

  const clearBox = async () => {
    let result = await Swal.fire({
      title: "Disconnect all parts from this box?",
      showDenyButton: true,
      confirmButtonText: "Yes",
      denyButtonText: `No`,
    })
    if (!result.isConfirmed) return

    for (const part of boxData.parts) {
      await api.updatePart(
        part.id,
        {
          name: part.name,
          description: part.description,
          properties: {},
          amount: part.amount,
          boxId: null,
          tags: part.tags?.map((i) => i.name) || [],
        },
        undefined
      )
    }

    await loadBoxData(scannedId)
  }

  let dataArea = <p>Please scan a box</p>
  if (scannedId != "No result") {
    dataArea = <p><span className='text-sm bg-gray-800'>{scannedId}</span> is not known.</p>
  }
  if (boxData) {
    const renderedParts = boxData.parts?.map(part => {
      return <div>
        <PartCard
          id={part.id}
          name={part.name}
          description={part.description}
          modifiedAt={new Date(part.updatedAt)}
          imageUrl={api.getImageUrl(part.imageId)}
          boxId={part.box?.id}
          positionId={part.box?.position?.id}
          tags={[]}
          properties={[]}
          actions={[
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
          checked={false}
          onCheckChanged={() => { }}
          onTagClick={() => { }}
        />
      </div>
    })

    dataArea = (
      <div>
        <p><span className='text-sm bg-gray-800'>{scannedId}</span> is known.</p>
        <p>Stored at: {boxData.position?.id || "not stored"}</p>
        <button onClick={clearBox} className="btn btn-error w-full mb-4">
          Clear Box
        </button>
        <div className='flex flex-col gap-2'>
          {renderedParts}
        </div>
      </div>
    )
  }

  return (
    <>
      <QrReader
        onResult={onResult}
        constraints={{ facingMode: { exact: "environment" } }}
        className="w-full"
        videoStyle={{ width: "100%", objectFit: "cover" }}
        videoContainerStyle={{ width: "100%" }}
      />
      <div className='p-2'>
        {dataArea}
      </div>
    </>
  );
};
