import React from "react"
import PartCard from "../components/partCard"

export default function Home() {
  return (
    <div className="h-full p-4">
      <div className="input-group w-full pb-4">
        <input type="text" placeholder="Search…" className="input input-bordered w-full" />
        <button className="btn btn-primary btn-square">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </button>
      </div>
      <main className="grid gap-2">
        <PartCard id={1} onActionClick={()=>{}} actionText="deliver" name="LUFTSTROM" />
        <PartCard id={2} onActionClick={()=>{}} actionText="deliver" name="EFFIZIENZ" />
        <PartCard id={3} onActionClick={()=>{}} actionText="deliver" name="WERBUNG" />
        <PartCard id={4} onActionClick={()=>{}} actionText="deliver" name="WOHN" />
        <PartCard id={5} onActionClick={()=>{}} actionText="deliver" name="WASCHMASCHINE" />
        <PartCard id={6} onActionClick={()=>{}} actionText="deliver" name="dymo band" />
        <PartCard id={7} onActionClick={()=>{}} actionText="deliver" name="y" />
        <PartCard id={8} onActionClick={()=>{}} actionText="deliver" name="feuerzeug" />
        <PartCard id={9} onActionClick={()=>{}} actionText="deliver" name="dw" />
      </main>
    </div>
  )
}
