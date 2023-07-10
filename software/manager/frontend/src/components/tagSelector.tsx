import React, { forwardRef } from "react"

import AsyncCreatableSelect from "react-select/async-creatable"
import { TagModel } from "../types"
import * as api from "../api"
import { ActionMeta } from "react-select"

type SelectOption = { label: string; value: string }

interface Props extends React.ComponentPropsWithoutRef<typeof AsyncCreatableSelect> {
  label?: string
  placeholder?: string
  error?: string
}

const transformTag = (tag: TagModel): SelectOption => {
  return {
    label: tag.name,
    value: tag.name,
  }
}

const filterTags = (inputValue: string, allTags: TagModel[]) => {
  return allTags.filter((i) => i.name.toLowerCase().includes(inputValue.toLowerCase())).map(transformTag)
}

const promiseOptions = (inputValue: string) =>
  new Promise<SelectOption[]>((resolve) => {
    api.getAllTags().then((tags) => {
      resolve(filterTags(inputValue, tags))
    })
  })

export default forwardRef(function TagSelector(props: Props, ref: React.ForwardedRef<AsyncCreatableSelect>) {
  const { label, placeholder, error, onChange, value, ...rest } = props

  const publicOnChange = (value: unknown, actionMeta: ActionMeta<unknown>) => {
    let selected = value as SelectOption[]
    onChange!(selected.map((v) => v.value).join(","), actionMeta)
  }

  const commaSeperatedToSelectOption = (value: string | undefined): SelectOption[] => {
    if (!value) return []
    return value.split(",").map((v) => ({ label: v, value: v }))
  }

  return (
    <AsyncCreatableSelect
      isMulti
      cacheOptions
      defaultOptions
      className="my-react-select-container"
      classNamePrefix="my-react-select"
      onChange={publicOnChange}
      value={commaSeperatedToSelectOption(value as string)}
      loadOptions={promiseOptions}
      {...rest}
    />
  )
})
