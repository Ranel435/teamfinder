export async function getAllHackathons() {
  const response = await fetch('/api/hackathons');
  if (!response.ok) throw new Error('Failed to fetch hackathons');
  return response.json();
}

export async function getProfilesByHackathonId(hackathonId: number) {
  const response = await fetch(`/api/hackathons/${hackathonId}/profiles`);
  if (!response.ok) throw new Error('Failed to fetch profiles');
  return response.json();
} 