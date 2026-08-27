import axios from "axios"

interface ShortUrlResponse {
  id: number,
  long_url: string,
  short_url: string
}

export const sendUrl = async (url: string) => {
  try {
    const res = await axios.post("http://localhost:4000/create-short-url", {
      long_url: url
    })

    if (!res.data) {
      throw new Error("Failed to create short URL");
    }

    return res.data as ShortUrlResponse
  } catch (error) {
    console.error("Error sending URL:", error)
    throw error
  }
}