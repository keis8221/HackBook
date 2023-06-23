import React from 'react'
import ItemsPage from './page'

type State = {
  items: {
    id: number
    name: string
  }[]
}

export default class ItemsContainer extends React.Component<{}, State> {
  constructor() {
    super({})
    this.state = {
      items: [],
    }
  }

  render() {
    return <ItemsPage items={this.state.items} />
  }
}
