import { v4 as uuidv4 } from 'uuid'
import { DraftRecord } from '@/models/record'
import type { ApiGetRecord } from '@/models/api/record'
import { mockUplaodableFile } from '@/mock/static/uploadableFile'

const fakeDescriptions = [
  'Matin du 31 octobre, après quelques chutes de neige sur les hauteurs',
  '31 octobre fin de matinée, après quelques kilomètres du bivouac',
  'Après midi du 31 octobre, le col des Cerces au fond',
  '31 octobre en fin de journée, passage au lac de la Clarée',
  'Arrivée à la voiture, fin de parcours',
  "Pas mal non ? C'est français.",
]

const fakeLocations = [
  'Lac des Béreaudes',
  'Col des Béreaudes',
  'Lac des Crouserocs',
  'Lac de la Ponsonnière',
  'Lac Blanc',
  'Pic du Galibier',
  'Col de la Ponsonnière',
  'Lac des Cerces',
]

const fakeDates = [
  '2024-10-30',
  '2024-10-31',
  '2024-11-10',
  '2024-11-15',
  '2024-12-20',
  '2025-01-02',
  '2025-03-16',
]

export const mockRecords = (length: number): Array<ApiGetRecord> => {
  const result: Array<ApiGetRecord> = []

  for (let i = 0; i < length; i++) {
    const uuid = uuidv4()
    const rndDescription = getRnd(fakeDescriptions.length)
    const rndLocation = getRnd(fakeLocations.length)
    const rndDates = getRnd(fakeDates.length)

    result.push({
      id: uuid,
      description: fakeDescriptions[rndDescription],
      location: fakeLocations[rndLocation],
      date: new Date(fakeDates[rndDates]).toISOString(),
    })
  }

  return result
}

export const mockDraftRecords = (length: number): Array<DraftRecord> => {
  const result: Array<DraftRecord> = []
  const files = mockUplaodableFile(length)

  for (let i = 0; i < length; i++) {
    result.push(new DraftRecord(files[i]))
  }

  return result
}

const getRnd = (max: number): number => {
  return Math.floor(Math.random() * max)
}
