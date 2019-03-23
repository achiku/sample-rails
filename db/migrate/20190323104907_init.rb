class Init < ActiveRecord::Migration[5.2]
  def self.up
    create_table :users do |t|
      t.text :name, null: false
      t.date :birthday, null: false
      t.timestamp :registered_at, null: false
    end

    create_table :todos do |t|
      t.references :users, foreign_key: true, index: true, null: false
      t.text :description, null:false
      t.boolean :is_done, null:false
      t.timestamp :registered_at, null:false
    end
  end

  def self.down
    drop_table :todos, force: :cascade
    drop_table :users, force: :cascade
  end
end
