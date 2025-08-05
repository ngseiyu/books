export interface Book {
  id: number
  title: string
  author: string
  genre: string
  description: string
  isbn: string
  published: string
  publisher: string
}

export const createEmptyBook = (): Book => ({
  id: 0,
  title: '',
  author: '',
  genre: '',
  description: '',
  isbn: '',
  published: '',
  publisher: '',
})
