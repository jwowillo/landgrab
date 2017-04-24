/// PlayerID identifies a Player by their position within a game.
enum PlayerID { noPlayer, player1, player2 }

/// Player of a game.
class Player {
  /// name of the Player.
  final String name;

  /// description of the Player.
  final String description;

  Map<String, dynamic> arguments = {};

  /// Player constructor initializes the Player's name and description.
  Player(this.name, {this.description: '', this.arguments}) {
    if (arguments == null) {
      arguments = {};
    }
  }
}
