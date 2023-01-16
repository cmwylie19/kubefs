import { render, screen } from '@testing-library/react';
import AppContainer from './AppContainer';

test('renders learn react link', () => {
  render(<AppContainer />);

  expect(screen.getByRole('app-container')).toHaveTextContent("")
});
