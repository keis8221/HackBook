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

  // async componentDidMount() {
  //   await getItems().then((items) => this.setState({ items }))
  // }

  render() {
    console.log(this.state.items)
    return <ItemsPage items={this.state.items} />
  }
}
