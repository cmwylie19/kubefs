import { render, fireEvent,screen } from '@testing-library/react';
import {Header,DisplayCount} from './Header';

test('renders learn react link', () => {
  render(<Header />);
  const linkElement = screen.getByText(/Kube-FS/i);
  expect(linkElement).toBeInTheDocument();
});

test('DisplayCount',  () => {
  let count = 0
  let result = DisplayCount((msg)=>"hi",count)

  expect(result).toBe("total: "+count)
})