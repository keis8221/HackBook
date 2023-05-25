'use client'
import { Button, TextField } from '@mui/material'
import { FormEventHandler } from 'react'

const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
  event.preventDefault()
  const { value: email } = (event.target as any).email
  const { value: password } = (event.target as any).password

  // TODO: send to server
  alert(`Email: ${email}\nPassword: ${password}`)
}

function SignUpPage() {
  return (
    <div className='flex justify-center mt-[8%]'>
      <div>
        <h1 className='text-3xl font-bold'>Sign Up</h1>
        <form className='mt-8' onSubmit={handleSubmit}>
          <label>
            <div>Email</div>
            <TextField
              id='email'
              label='Email'
              type='email'
              sx={{ m: 1, width: '50ch' }}
              placeholder='example@example.com'
              hiddenLabel={true}
            />
          </label>

          <label className='mt-4'>
            <div>Password</div>
            <TextField
              id='password'
              label='Password'
              type='password'
              sx={{ m: 1, width: '50ch' }}
              hiddenLabel
            />
          </label>

          <label className='mt-4'>
            <div>ReEnter-Password</div>
            <TextField
              id='reEnterPassword'
              label='ReEnter-Password'
              type='password'
              sx={{ m: 1, width: '50ch' }}
              hiddenLabel
            />
          </label>

          <Button variant='outlined' type='submit' className='mt-4 block'>
            Next
          </Button>
        </form>
        <a href='/login' className='mt-4 text-[#509ef6]'>
          Already have an account? Log in
        </a>
      </div>
    </div>
  )
}

export default SignUpPage
