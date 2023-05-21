'use client'
import React from 'react'

type ItemsPageProps = {
  items: {
    id: number
    name: string
  }[]
}

export const ItemsPage = ({ items }: ItemsPageProps) => {
  return (
    <div className='flex justify-center'>
      <h1>あなたのおすすめ</h1>
      {items?.map((item) => (
        <div key={item.id}>{item.name}</div>
      ))}
    </div>
  )
}
export default ItemsPage
