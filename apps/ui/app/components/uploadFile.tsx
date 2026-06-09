'use client'

import { ChangeEvent, useState } from "react";

export default function UploadGpxForm() {

  const [file, setFile] = useState<File | null>(null)

  const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setFile(e.target.files[0])
    }
  }

  const handleFileUpload = async () => {
    const url = "http://localhost:8080/gpx"
    if (!file) {
      return
    }

    const formData = new FormData();
    formData.append("file", file)

    const response = await fetch(url, {
      method: "POST",
      body: formData
    })

    console.log(response.status)
  }

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleFileUpload}>Upload</button>

      {file && (
        <dl>
          <dt>Filename:</dt>
          <dd>{file.name}</dd>

          <dt>Filesize:</dt>
          <dd>{(file.size / 1024).toFixed(2)}KB</dd>
        </dl>
      )}

    </div>
  );
}
