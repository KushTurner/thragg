import { describe, it, expect, beforeEach } from 'vitest'

// ============================================================
// ARRANGE / ACT / ASSERT — one clear outcome per test
// ============================================================

describe('UserService', () => {
  let userService: UserService
  let db: Database

  beforeEach(async () => {
    // Each test gets a fresh DB — no shared state
    db = await createTestDatabase()
    userService = new UserService(db)
  })

  it('returns the created user', async () => {
    // Arrange
    const input = { name: 'Alice', email: 'alice@example.com' }

    // Act
    const user = await userService.createUser(input)

    // Assert
    expect(user.name).toBe('Alice')
    expect(user.email).toBe('alice@example.com')
    expect(user.id).toBeDefined()
  })

  // ============================================================
  // TEST ERROR PATHS — not just the happy path
  // ============================================================

  it('throws on duplicate email', async () => {
    // Arrange
    const input = { name: 'Alice', email: 'alice@example.com' }
    await userService.createUser(input)

    // Act & Assert
    await expect(userService.createUser(input)).rejects.toThrow('duplicate email')
  })

  it('throws on missing email', async () => {
    // Arrange
    const input = { name: 'Alice', email: '' }

    // Act & Assert
    await expect(userService.createUser(input)).rejects.toThrow('email is required')
  })
})

// ============================================================
// REAL DATABASE — not mocks or in-memory substitutes
// Use testcontainers-node for integration tests
// ============================================================

async function createTestDatabase(): Promise<Database> {
  const { PostgreSqlContainer } = await import('@testcontainers/postgresql')

  const container = await new PostgreSqlContainer('postgres:16').start()

  const db = new Database(container.getConnectionUri())
  await db.migrate()

  return db
}
