import { useEffect, useState } from 'react'
import './App.css'

const API_BASE = 'http://localhost:8000/api/v2'

function App() {
  const [recordId, setRecordId] = useState('1')
  const [versions, setVersions] = useState([])
  const [selectedVersion, setSelectedVersion] = useState(null)

  const [fromVersion, setFromVersion] = useState('')
  const [toVersion, setToVersion] = useState('')

  const [analysis, setAnalysis] = useState('')
  const [loading, setLoading] = useState(false)

  // dynamic underwriting form
  const [updateForm, setUpdateForm] = useState({})

  async function loadRecord() {
    try {
      await loadVersions()
    } catch (err) {
      console.error(err)
    }
  }

  async function loadVersions() {
    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}/versions`
      )

      const data = await response.json()

      const sorted = [...data].sort(
        (a, b) => new Date(b.created_at) - new Date(a.created_at)
      )

      setVersions(sorted)

    } catch (err) {
      console.error(err)
    }
  }

  async function loadVersion(version) {
    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}/versions/${version}`
      )

      const data = await response.json()

      setSelectedVersion(data)

      // dynamically populate ALL fields
      setUpdateForm(data.data || {})

    } catch (err) {
      console.error(err)
    }
  }

  async function analyzeVersions() {

    if (!fromVersion || !toVersion) {
      alert('Please select both dates')
      return
    }

    setLoading(true)
    setAnalysis('')

    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}/analyze`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            from_version: Number(fromVersion),
            to_version: Number(toVersion),
          }),
        }
      )

      const data = await response.json()

      setAnalysis(
        data.analysis || JSON.stringify(data, null, 2)
      )

    } catch (err) {
      console.error(err)
      setAnalysis('Failed to analyze versions')
    }

    setLoading(false)
  }

  async function updateRecord() {

    if (!selectedVersion) {
      alert('Select a historical version first')
      return
    }

    try {
      const response = await fetch(
        `${API_BASE}/records/${recordId}`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(updateForm),
        }
      )

      if (!response.ok) {
        throw new Error('Failed to update record')
      }

      await loadRecord()

      alert('New version created successfully')

    } catch (err) {
      console.error(err)
      alert('Failed to update record')
    }
  }

  useEffect(() => {
    loadRecord()
  }, [])

  const panelStyle = {
    background: '#1e293b',
    borderRadius: '18px',
    padding: '24px',
    boxShadow: '0 10px 30px rgba(0,0,0,0.25)',
  }

  const inputStyle = {
    width: '100%',
    padding: '12px',
    borderRadius: '10px',
    border: '1px solid #334155',
    background: '#0f172a',
    color: 'white',
    boxSizing: 'border-box',
  }

  return (
    <div
      style={{
        minHeight: '100vh',
        background: '#020617',
        color: 'white',
        padding: '32px',
        fontFamily: 'Inter, sans-serif',
      }}
    >
      <div
        style={{
          maxWidth: '1500px',
          margin: '0 auto',
        }}
      >

        {/* HEADER */}
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            marginBottom: '32px',
          }}
        >
          <div>
            <h1
              style={{
                fontSize: '42px',
                marginBottom: '8px',
              }}
            >
              TimeTravel Underwriting Platform
            </h1>

            <p
              style={{
                color: '#94a3b8',
              }}
            >
              Temporal business reconstruction and underwriting intelligence.
            </p>
          </div>

          <div
            style={{
              display: 'flex',
              gap: '12px',
            }}
          >
            <input
              value={recordId}
              onChange={(e) =>
                setRecordId(e.target.value)
              }
              placeholder="Record ID"
              style={{
                ...inputStyle,
                width: '160px',
              }}
            />

            <button
              onClick={loadRecord}
              style={{
                background: '#2563eb',
                color: 'white',
                border: 'none',
                borderRadius: '10px',
                padding: '12px 20px',
                cursor: 'pointer',
                fontWeight: 'bold',
              }}
            >
              Load Record
            </button>
          </div>
        </div>

        {/* GRID */}
        <div
          style={{
            display: 'grid',
            gridTemplateColumns: '1fr 1fr',
            gap: '24px',
          }}
        >

          {/* TIMELINE */}
          <div style={panelStyle}>
            <h2 style={{ marginBottom: '20px' }}>
              Historical Timeline
            </h2>

            <div
              style={{
                maxHeight: '500px',
                overflowY: 'auto',
              }}
            >
              {versions.map((v) => (
                <div
                  key={`${v.version}-${v.created_at}`}
                  onClick={() => loadVersion(v.version)}
                  style={{
                    background:
                      selectedVersion?.created_at === v.created_at
                        ? '#2563eb'
                        : '#334155',

                    padding: '16px',
                    borderRadius: '12px',
                    marginBottom: '12px',
                    cursor: 'pointer',
                    transition: '0.2s',
                  }}
                >
                  <div
                    style={{
                      fontWeight: 'bold',
                      marginBottom: '6px',
                    }}
                  >
                    {new Date(v.created_at).toLocaleString()}
                  </div>

                  <div
                    style={{
                      color: '#cbd5e1',
                      fontSize: '14px',
                    }}
                  >
                    Snapshot Version {v.version}
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* SNAPSHOT VIEWER */}
          <div style={panelStyle}>
            <h2 style={{ marginBottom: '20px' }}>
              Snapshot Viewer
            </h2>

            {selectedVersion ? (
              <pre
                style={{
                  whiteSpace: 'pre-wrap',
                  color: '#e2e8f0',
                  lineHeight: '1.6',
                }}
              >
                {JSON.stringify(
                  selectedVersion,
                  null,
                  2
                )}
              </pre>
            ) : (
              <p style={{ color: '#94a3b8' }}>
                Select a historical version to inspect or edit.
              </p>
            )}
          </div>

          {/* UPDATE RECORD */}
          <div style={panelStyle}>
            <h2 style={{ marginBottom: '20px' }}>
              Update Existing Record
            </h2>

            {selectedVersion ? (
              <>
                <div
                  style={{
                    background: '#0f172a',
                    padding: '16px',
                    borderRadius: '12px',
                    marginBottom: '20px',
                    border: '1px solid #334155',
                  }}
                >
                  <div
                    style={{
                      fontSize: '20px',
                      fontWeight: 'bold',
                      marginBottom: '8px',
                    }}
                  >
                    Editing Snapshot From Timeline
                  </div>

                  <div
                    style={{
                      color: '#94a3b8',
                      fontSize: '14px',
                      marginBottom: '8px',
                    }}
                  >
                    Selected snapshot becomes the base for this update.
                  </div>

                  <div
                    style={{
                      color: '#cbd5e1',
                    }}
                  >
                    Record ID: {recordId}
                  </div>
                </div>

                {/* DYNAMIC FORM */}
                <div
                  style={{
                    display: 'grid',
                    gap: '14px',
                  }}
                >
                  {Object.entries(updateForm).map(
                    ([key, value]) => (
                      <div key={key}>
                        <div
                          style={{
                            marginBottom: '6px',
                            color: '#cbd5e1',
                            fontSize: '14px',
                            textTransform: 'capitalize',
                          }}
                        >
                          {key.replaceAll('_', ' ')}
                        </div>

                        <input
                          value={value}
                          onChange={(e) =>
                            setUpdateForm({
                              ...updateForm,
                              [key]: e.target.value,
                            })
                          }
                          style={inputStyle}
                        />
                      </div>
                    )
                  )}

                  <button
                    onClick={updateRecord}
                    style={{
                      background: '#16a34a',
                      color: 'white',
                      border: 'none',
                      borderRadius: '10px',
                      padding: '14px',
                      cursor: 'pointer',
                      fontWeight: 'bold',
                      marginTop: '8px',
                    }}
                  >
                    Create New Version
                  </button>
                </div>
              </>
            ) : (
              <p style={{ color: '#94a3b8' }}>
                Select a historical version from the timeline
                to begin editing.
              </p>
            )}
          </div>

          {/* ANALYSIS */}
          <div style={panelStyle}>
            <h2 style={{ marginBottom: '20px' }}>
              Underwriting Analysis
            </h2>

            <div
              style={{
                display: 'flex',
                flexDirection: 'column',
                gap: '12px',
                marginBottom: '20px',
              }}
            >
              <select
                value={fromVersion}
                onChange={(e) =>
                  setFromVersion(e.target.value)
                }
                style={inputStyle}
              >
                <option value="">
                  Select From Date
                </option>

                {versions.map((v) => (
                  <option
                    key={`from-${v.version}`}
                    value={v.version}
                  >
                    {new Date(
                      v.created_at
                    ).toLocaleString()}
                  </option>
                ))}
              </select>

              <select
                value={toVersion}
                onChange={(e) =>
                  setToVersion(e.target.value)
                }
                style={inputStyle}
              >
                <option value="">
                  Select To Date
                </option>

                {versions.map((v) => (
                  <option
                    key={`to-${v.version}`}
                    value={v.version}
                  >
                    {new Date(
                      v.created_at
                    ).toLocaleString()}
                  </option>
                ))}
              </select>

              <button
                onClick={analyzeVersions}
                disabled={loading}
                style={{
                  background: '#7c3aed',
                  color: 'white',
                  border: 'none',
                  borderRadius: '10px',
                  padding: '14px',
                  cursor: 'pointer',
                  fontWeight: 'bold',
                }}
              >
                {loading
                  ? 'Analyzing...'
                  : 'Analyze Changes'}
              </button>
            </div>

            <div
              style={{
                background: '#0f172a',
                borderRadius: '12px',
                padding: '20px',
                minHeight: '250px',
              }}
            >
              <pre
                style={{
                  whiteSpace: 'pre-wrap',
                  color: '#e2e8f0',
                  lineHeight: '1.6',
                }}
              >
                {analysis ||
                  'No underwriting analysis generated yet.'}
              </pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App