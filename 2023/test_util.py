import pytest
import tempfile
from utils import read_file_chunks

@pytest.mark.asyncio
async def test_read_file_chunks():
  with tempfile.TemporaryDirectory() as tempdir:
    file = f"{tempdir}/test.txt"
    with open(file, "w") as f:
      f.write("a\nb\n\nc\nd\n\ne\nf\n\ng\nh\n\n")
      
    expected = [
      ["a", "b"],
      ["c", "d"],
      ["e", "f"],
      ["g", "h"],
    ]
    i = 0
    async for line in read_file_chunks(file, 2):
      assert len(line) == 2
      assert line[0] != ""
      assert line[1] != ""

      assert line == expected[i]
      i += 1
      
@pytest.mark.asyncio
async def test_read_file_chunks_longer():
  with tempfile.TemporaryDirectory() as tempdir:
    file = f"{tempdir}/test.txt"
    expected = [
      ["Brazil", "Brasilia", "Portuguese", "South America"],
      ["Argentina", "Buenos Aires", "Spanish", "South America"],
      ["Uruguay", "Montevideo", "Spanish", "South America"],
      ["Paraguay", "Asuncion", "Spanish", "South America"],
      ["France", "Paris", "French", "Europe"],
      ["Spain", "Madrid", "Spanish", "Europe"],
      ["Portugal", "Lisbon", "Portuguese", "Europe"],
      ["Italy", "Rome", "Italian", "Europe"],
      ["United Kingdom", "London", "English", "Europe"],
      ["United States", "Washington", "English", "North America"],
      ["Nigeria", "Abuja", "English", "Africa"],
      ["China", "Beijing", "Mandarin", "Asia"],
      ["Japan", "Tokyo", "Japanese", "Asia"],
      ["India", "New Delhi", "Hindi", "Asia"],
    ]
    
    with open(file, "w") as f:
      f.write("Brazil\nBrasilia\nPortuguese\nSouth America\n\n")
      f.write("Argentina\nBuenos Aires\nSpanish\nSouth America\n\n")
      f.write("Uruguay\nMontevideo\nSpanish\nSouth America\n\n")
      f.write("Paraguay\nAsuncion\nSpanish\nSouth America\n\n")
      f.write("France\nParis\nFrench\nEurope\n\n")
      f.write("Spain\nMadrid\nSpanish\nEurope\n\n")
      f.write("Portugal\nLisbon\nPortuguese\nEurope\n\n")
      f.write("Italy\nRome\nItalian\nEurope\n\n")
      f.write("United Kingdom\nLondon\nEnglish\nEurope\n\n")
      f.write("United States\nWashington\nEnglish\nNorth America\n\n")
      f.write("Nigeria\nAbuja\nEnglish\nAfrica\n\n")
      f.write("China\nBeijing\nMandarin\nAsia\n\n")
      f.write("Japan\nTokyo\nJapanese\nAsia\n\n")
      f.write("India\nNew Delhi\nHindi\nAsia\n\n")
      
    i = 0
    async for line in read_file_chunks(file, 4):
      assert len(line) == 4
      assert line[0] != ""
      assert line[1] != ""

      assert line == expected[i]
      i += 1