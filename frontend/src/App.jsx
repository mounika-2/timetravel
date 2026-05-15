import { useState } from 'react'

const API_BASE = 'http://localhost:8000/api/v2'

export default function App() {
  const [recordId, setRecordId] = useState('1')
  const [versions, setVersions] = useState([])
  const [selectedVersion, setSelectedVersion] = useState(null)
  const [updateJSON, setUpdateJSON] = useState('{\n  "employees": "100"\n}')
  const [error, setError] = useState('')

  async function loadVersions() {
    setError('')

    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}/versions`
      )

      if (!response.ok) {
        throw new Error('Could not load versions')
      }

      const data = await response.json()

      setVersions(data)

      if (data.length > 0) {
        setSelectedVersion(data[data.length - 1])
      }
    } catch (err) {
      setError(err.message)
    }
  }

    async function loadSpecificVersion(version) {
    setError('')

    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}/versions/${version}`
      )

      if (!response.ok) {
        throw new Error('Could not load version')
      }

      const data = await response.json()

      setSelectedVersion(data)
    } catch (err) {
      setError(err.message)
    }
  }

    async function submitUpdate() {
    setError('')

    try {
      const parsed = JSON.parse(updateJSON)

      const response = await fetch(
        `http://localhost:8000/api/v1/records/${recordId}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(parsed),
        }
      )

      if (!response.ok) {
        throw new Error('Update failed')
      }

      await loadVersions()
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div className="min-h-screen p-8 bg-gray-100">
      <div className="max-w-6xl mx-auto">
        <h1 className="text-4xl font-bold mb-6">
          Time Travel Records
        </h1>

        <div className="bg-white rounded-xl shadow p-4 mb-6">
          <div className="flex gap-4 items-center">
            <input
              type="text"
              value={recordId}
              onChange={(e) => setRecordId(e.target.value)}
              className="border rounded px-4 py-2 w-48"
              placeholder="Record ID"
            />

            <button
              onClick={loadVersions}
              className="bg-black text-white px-4 py-2 rounded"
            >
              Load Record
            </button>
          </div>
        </div>

        {error && (
          <div className="bg-red-100 border border-red-300 text-red-700 p-4 rounded mb-6">
            {error}
          </div>
        )}

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="bg-white rounded-xl shadow p-4">
            <h2 className="text-xl font-semibold mb-4">
              Versions
            </h2>

            <div className="space-y-2">
              {versions.map((version) => (
                <button
                  key={version.version}
                  onClick={() => loadSpecificVersion(version.version)}
                  className="w-full text-left border rounded p-3 hover:bg-gray-100"
                >
                  <div className="font-semibold">
                    Version {version.version}
                  </div>

                  <div className="text-sm text-gray-500">
                    {new Date(version.created_at).toLocaleString()}
                  </div>
                </button>
              ))}
            </div>
          </div>
         <div className="bg-white rounded-xl shadow p-4 md:col-span-2">
            <h2 className="text-xl font-semibold mb-4">
              Record Snapshot
            </h2>

            {selectedVersion ? (
              <div>
                <div className="mb-4">
                  <div className="font-semibold">
                    Version {selectedVersion.version}
                  </div>

                  <div className="text-sm text-gray-500">
                    {new Date(selectedVersion.created_at).toLocaleString()}
                  </div>
                </div>

                <pre className="bg-gray-100 rounded p-4 overflow-auto text-sm">
                  {JSON.stringify(selectedVersion.data, null, 2)}
                </pre>
              </div>
            ) : (
              <div>No version selected</div>
            )}
          </div>
        </div>
          <div className="bg-white rounded-xl shadow p-4 mt-6">
          <h2 className="text-xl font-semibold mb-4">
            Update Record
          </h2>

          <textarea
            value={updateJSON}
            onChange={(e) => setUpdateJSON(e.target.value)}
            className="w-full h-48 border rounded p-4 font-mono text-sm"
          />

          <button
            onClick={submitUpdate}
            className="mt-4 bg-black text-white px-4 py-2 rounded"
          >
            Submit Update
          </button>
        </div>
      </div>
    </div>
  )
}