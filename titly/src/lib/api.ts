import axios from "axios"

export const sendUrl = async (url: string) => {
  try {
    const res = await axios.post("http://localhost:4000/create-short-url", {
      long_url: url
    })
    return res.data
  } catch (error) {
    console.error("Error sending URL:", error)
  }
}