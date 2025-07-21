import os
import sys
import requests
import json
from typing import List, Dict, Any
import argparse

class BookSearchClient:
    def __init__(self, api_host: str, api_port: str):
        self.base_url = f"http://{api_host}:{api_port}"
        self.session = requests.Session()
        self.session.timeout = 30

    def search_books(self, query: str) -> Dict[str, Any]:
        """Search for books using the API."""
        url = f"{self.base_url}/search"
        params = {'q': query}
        
        print(f"Searching for: '{query}'...")
        
        try:
            response = self.session.get(url, params=params)
            response.raise_for_status()
            return response.json()
        except requests.exceptions.ConnectionError:
            raise Exception(f"Could not connect to API at {self.base_url}")
        except requests.exceptions.RequestException as e:
            raise Exception(f"Request failed: {e}")

    def format_book(self, book: Dict[str, Any]) -> str:
        """Format a single book for display."""
        title = book.get('title', 'Unknown Title')
        authors = book.get('author_name', [])
        year = book.get('first_publish_year', '')
        
        author_str = ', '.join(authors[:3]) if authors else 'Unknown Author'
        if len(authors) > 3:
            author_str += f" (and {len(authors) - 3} more)"
        
        year_str = f" ({year})" if year else ""
        
        return f"{title}\n   {author_str}{year_str}"

    def display_results(self, data: Dict[str, Any]) -> None:
        """Display search results in a human-readable format."""
        total = data.get('total', 0)
        books = data.get('books', [])
        
        print(f"\n{'='*60}")
        print(f"Found {total} results")
        print(f"{'='*60}")
        
        if not books:
            print("No books found for your search.")
            return
        
        for i, book in enumerate(books, 1):
            print(f"\n{i:2d}. {self.format_book(book)}")
        
        print(f"\n{'='*60}")
        print(f"Showing {len(books)} results out of {total} total")
        print(f"{'='*60}")

def main():
    parser = argparse.ArgumentParser(description='Search for books')
    parser.add_argument('query', help='Search query')
    parser.add_argument('--host', default=os.getenv('API_HOST', 'localhost'), 
                       help='API host (default: localhost)')
    parser.add_argument('--port', default=os.getenv('API_PORT', '3000'), 
                       help='API port (default: 3000)')
    
    args = parser.parse_args()
    
    try:
        client = BookSearchClient(args.host, args.port)
        data = client.search_books(args.query)
        client.display_results(data)
    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()