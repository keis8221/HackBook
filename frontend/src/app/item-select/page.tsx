'use client'
import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import CardHeader from '@mui/material/CardHeader'
import CardMedia from '@mui/material/CardMedia'
import Radio from '@mui/material/Radio'
import Typography from '@mui/material/Typography'
import Image from 'next/image'
import React, { Suspense } from 'react'
import { getItems } from '../api/item'

type Image = {
  id: number
  url: string
}

const demoImageData: Image[] = [
  {
    id: 1,
    url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTl36Cp-tnywJtenfEggyc0gGGx3s0zUH_Y2g&usqp=CAU',
  },
  {
    id: 2,
    url: 'https://images.pexels.com/photos/4386158/pexels-photo-4386158.jpeg?auto=compress&cs=tinysrgb&w=800g',
  },
  {
    id: 3,
    url: 'https://baseec-img-mng.akamaized.net/images/item/origin/1902cd1aa84b637fdd87d5970aceed39.jpg?imformat=generic',
  },
  {
    id: 4,
    url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRCKBDMqmC-oodlw3xLOD9Wnt6B3sCmG0pCfw&usqp=CAU',
  },
  {
    id: 5,
    url: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTZriGaoERHOAEzHWpzs60i3D2rOceDNOsrnw&usqp=CAU',
  },
]

const ImageSelectionForm: React.FC = () => {
  const [selectedImage, setSelectedImage] = React.useState<Image | null>(null)
  // const router = useRouter()

  const images: Image[] = demoImageData

  const handleImageSelect = (image: Image) => {
    setSelectedImage(image)
  }

  const [item, setItems] = React.useState<any>()
  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const newItem = await getItems()
    setItems(newItem)
  }

  return (
    <form className='m-4' onSubmit={handleSubmit}>
      <h1 className='text-3xl font-bold text-center'>Choose an image of your interest!!</h1>
      {images.map((image) => (
        <div key={image.id} className='flex justify-center mt-5'>
          <label className='relative'>
            <Radio
              sx={{
                '& .MuiSvgIcon-root': {
                  fontSize: 28,
                },
              }}
              value={image.id}
              checked={selectedImage?.id === image.id}
              onChange={() => handleImageSelect(image)}
              className={`absolute top-[50%] left-[320px] z-10`}
            />
            <Suspense fallback={<div>Loading...</div>}>
              <img src={image.url} alt='' width={300} height={300} className='rounded-xl' />
            </Suspense>
          </label>
        </div>
      ))}
      <div className='flex justify-center mt-5'>
        <Button variant='outlined' type='submit'>
          Submit
        </Button>
      </div>

      {item ? (
        <div className='flex justify-center mt-5'>
          <Card sx={{ maxWidth: 345 }}>
            <CardHeader title={item.Items[0].Item.title} subheader={item.Items[0].Item.author} />
            <CardMedia component='img' image={item.Items[0].Item.largeImageUrl} alt='' />
            <CardContent>
              <Typography variant='body2' color='text.secondary'>
                {item.Items[0].Item.itemCaption}
              </Typography>
              <a className='text-[#52d1f8]' href={item.Items[0].Item.itemUrl}>
                go to product Page
              </a>
            </CardContent>
          </Card>
        </div>
      ) : (
        ''
      )}
    </form>
  )
}

export default ImageSelectionForm
