# Changelog

## [2.0.0.beta2] - 2022-04-16

### Added
- A new table sorting menu is shown on small device screens, allowing users to
  sort tables in both directions by columns even when they're hidden or their
  information is not displayed as a table column directly.

### Changed
- The information from hidden table columns on small device screens now shows
  up directly below the item's name. Smartphone users are no longer at a
  disadvantage when it comes to the amount of information shown to them.
- All frontend code that deals with account management was moved from Go-based
  templates to the Vue-based login app.

### Removed
- Password confirmation fields have been removed entirely. They've been
  replaced with a button on the regular password field that toggles the
  password's visibility.

### Fixed
- It is no longer possible for sliders to move beyond their bounds by entering
  very small or very large numbers in the associated input fields.

[2.0.0.beta2]: https://github.com/and-rad/heyapple/compare/v2.0.0.beta1...v2.0.0.beta2